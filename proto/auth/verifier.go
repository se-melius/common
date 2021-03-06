package authpb

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/zoenion/common/errors"
)

type TokenVerifier interface {
	Verify(ctx context.Context, t *JWT) (JWTState, error)
}

type tokenVerifier struct {
	sync.Mutex
	key *ecdsa.PublicKey
}

func (v *tokenVerifier) verifySignature(t *JWT) (bool, error) {
	if t.Claims == nil {
		return false, errors.Forbidden
	}

	claimsBytes, err := proto.Marshal(t.Claims)
	if err != nil {
		return false, fmt.Errorf("could not encode claims: %s", err)
	}

	sha := sha256.New()
	sha.Write(claimsBytes)
	hash := sha.Sum(nil)

	parts := strings.Split(t.Signature, ".")
	r, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, errors.New("token wrong format")
	}
	s, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, errors.New("token wrong format")
	}

	rInt := new(big.Int)
	rInt.SetBytes(r)
	sInt := new(big.Int)
	sInt.SetBytes(s)

	if t.Header.Alg == "ecdsa" {
		return ecdsa.Verify(v.key, hash, rInt, sInt), nil

	} else {
		return false, errors.NotSupported
	}
}

func (v *tokenVerifier) verifyToken(ctx context.Context, t *JWT) (JWTState, error) {
	verified, err := v.verifySignature(t)
	if err != nil {
		return 0, err
	}

	if !verified {
		return JWTState_NOT_SIGNED, nil
	}

	if t.Claims.Exp != -1 && t.Claims.Exp <= time.Now().Unix() {
		return JWTState_EXPIRED, nil
	}

	if t.Claims.Nbf != -1 && t.Claims.Nbf > time.Now().Unix() {
		return JWTState_NOT_EFFECTIVE, nil
	}

	return JWTState_VALID, nil
}

func (v *tokenVerifier) Verify(ctx context.Context, t *JWT) (JWTState, error) {
	if t == nil {
		return JWTState_NOT_VALID, errors.Forbidden
	}

	state, err := v.verifyToken(ctx, t)
	if err != nil {
		return JWTState_NOT_VALID, errors.Forbidden
	}
	return state, nil
}

func NewTokenVerifier(key *ecdsa.PublicKey) *tokenVerifier {
	return &tokenVerifier{
		key: key,
	}
}

func TokenFromJWT(jwt string) (*JWT, error) {
	if jwt == "" {
		return nil, nil
	}

	jwt = strings.Replace(jwt, "Bearer ", "", 1)

	malformed := errors.New("malformed token")
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, errors.New("missing parts")
	}

	var t JWT
	t.Header = new(JWTHeader)
	t.Claims = new(Claims)

	headerBytes, _ := base64.StdEncoding.DecodeString(parts[0])
	if headerBytes == nil {
		return nil, malformed
	}

	claimsBytes, _ := base64.StdEncoding.DecodeString(parts[1])
	if claimsBytes == nil {
		return nil, malformed
	}

	signatureBytes, _ := base64.StdEncoding.DecodeString(parts[2])
	if signatureBytes == nil {
		return nil, malformed
	}

	err := json.Unmarshal(headerBytes, t.Header)
	if err != nil {
		return nil, malformed
	}

	err = json.Unmarshal(claimsBytes, t.Claims)
	if err != nil {
		return nil, malformed
	}

	err = json.Unmarshal(signatureBytes, &t.Signature)
	if err != nil {
		return nil, malformed
	}

	return &t, nil
}

type StringTokenVerifier struct {
	verifier TokenVerifier
}

func (stv *StringTokenVerifier) Verify(ctx context.Context, jwt string) (context.Context, error) {
	t, err := TokenFromJWT(jwt)
	if err != nil {
		return ctx, err
	}
	_, err = stv.verifier.Verify(ctx, t)
	return ctx, err
}

func NewStringTokenVerifier(tv TokenVerifier) *StringTokenVerifier {
	return &StringTokenVerifier{
		verifier: tv,
	}
}

func String(jwt *JWT) (string, error) {
	headerBytes, err := json.Marshal(jwt.Header)
	if err != nil {
		return "", err
	}

	claimsBytes, err := json.Marshal(jwt.Claims)
	if err != nil {
		return "", err
	}

	signatureBytes, err := json.Marshal(jwt.Signature)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s.%s",
		base64.StdEncoding.EncodeToString(headerBytes),
		base64.StdEncoding.EncodeToString(claimsBytes),
		base64.StdEncoding.EncodeToString(signatureBytes),
	), nil
}
