package pkce

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkce_ChallengeCode(t *testing.T) {
	p := Pkce{
		RandomString: "45d7820e694481f399e7fb9c444f0cb63301a7254d1401443835d9af2c9a6a5ec5b243c3470feb945336025964ef05c8d2f0e44baf76762ba6136914",
	}

	assert.Equal(t, "iQoF8w9kq5RnuMdisRXypyOoMCF7FGz-ro7dwHjC28U", p.ChallengeCode())
}

func TestPkce_VerifyCode(t *testing.T) {
	p := Pkce{
		RandomString: "ThisIsAVeryBiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiigRandomString",
	}

	code, err := p.VerifyCode()

	if assert.Nil(t, err) {
		assert.Equal(t, "ThisIsAVeryBiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiigRandomString", code)
	}
}

func TestPkce_VerifyCode2(t *testing.T) {
	p := Pkce{
		Length: 50,
	}

	code, err := p.VerifyCode()

	if assert.Nil(t, err) {
		assert.Equal(t, len(code), 50)
	}
}

func TestPkce_VerifyCode3(t *testing.T) {
	p := Pkce{}

	_, err := p.VerifyCode()

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("length should be greater than 0"), err)
	}
}

func TestPkce_VerifyCode4(t *testing.T) {
	p := Pkce{
		Length: 42,
	}

	_, err := p.VerifyCode()

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("length should be >=43 and <=128"), err)
	}
}

func TestPkce_VerifyCode5(t *testing.T) {
	p := Pkce{
		RandomString: "ThisIsSmallRandomString",
	}

	_, err := p.VerifyCode()

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("length should be >=43 and <=128"), err)
	}
}

func ExamplePkce_VerifyCode() {
	p := Pkce{
		RandomString: "ThisIsAVeryBiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiigRandomString",
	}

	code, err := p.VerifyCode()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
	// Output: ThisIsAVeryBiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiigRandomString
}

func ExamplePkce_VerifyCode_generatedString() {
	p := Pkce{
		Length: 60,
	}

	code, err := p.VerifyCode()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}

func ExamplePkce_ChallengeCode() {
	p := Pkce{
		RandomString: "45d7820e694481f399e7fb9c444f0cb63301a7254d1401443835d9af2c9a6a5ec5b243c3470feb945336025964ef05c8d2f0e44baf76762ba6136914",
	}
	fmt.Println(p.ChallengeCode())
	// Output: iQoF8w9kq5RnuMdisRXypyOoMCF7FGz-ro7dwHjC28U
}

func ExamplePkce_ChallengeCode_generatedString() {
	p := Pkce{
		Length: 60,
	}
	code, err := p.VerifyCode()
	if err != nil {
		panic(err)
	}
	fmt.Println(code, p.ChallengeCode())
}
