package jwt

import (
	"gf_demo_api/app/jsonapi"
	"testing"

	"github.com/gogf/gf/util/guid"
)

func TestCreateToken(t *testing.T) {
	strToken, err := CreateToken(jsonapi.Token{
		Uin:  10000,
		Skey: guid.S(),
	})
	t.Log(strToken, err)
}

func TestParseToken(t *testing.T) {
	rsp, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTA2NjA5MDUsImp0aSI6ImI2YjkwMmI3LTMxNWItNDY1ZS1iNmIzLTEzMDFjZmI1ODU5YSIsImlzcyI6ImRlbW8uY29tIiwiSnd0U2Vzc2lvbiI6eyJzb3VyY2UiOiIiLCJpcCI6IiIsIm1rIjoiIiwidWluIjoxMDAwMCwicm9sZSI6IiIsInNrZXkiOiI3NjQ2OTlhZS0zYTQ5LTRiODItOTU0OS03ODBkNDVjZWFmYTkifX0.2w1PvimCANAE38R951zC5fNYEKStZ7xWwsj55cWz5qo")
	t.Log(rsp, err)
}
