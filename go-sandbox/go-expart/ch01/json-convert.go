// 構造体の方をエクスポートせずにフィールドをエクスポートする例
// encodeing/json
package vault

import (
	"bytes"
	"io"
	"time"
)

type secret struct {
	ID string
	CreateTime time.Time

	token string
}

func(s *secret) Read(p []byte) (int, error) {
	return bytes.NewBuffer(p).WriteString(s.token)
}

func NewSecret() io.Reader {
	return &secret{
		ID: "dummy_id",
		CreateTime: time.Now(),
		token: "dummy_token",
	}
}

s := vault.NewSecret()

err := json.NewEncoder(os.Stdout).Encode(s)
if err != nil {
	fmt.Println("failed to json encode, error=", err)
}
