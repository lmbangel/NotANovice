package agents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Agent interface {
	Prompt() Response
}

type Request struct {
	Model  string
	Prompt string
	//UserPrompt   string
}

type Response struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
}

type Ollama struct {
	Url     string   `json:"url"`
	Request *Request `json:"request"`
}

func (o *Ollama) Prompt() *Response {

	data, _ := json.Marshal(o.Request)

	res, err := http.Post(o.Url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return &Response{
			Model:     o.Request.Model,
			CreatedAt: time.Now(),
			Response:  fmt.Sprintf("Error: %s", err.Error()),
			Done:      true,
		}
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	lines := bytes.Split(body, []byte("\n"))

	var output string
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var res Response
		if err := json.Unmarshal(line, &res); err == nil {
			output += res.Response
			if res.Done {
				break
			}
		}
	}

	return &Response{
		Model:     o.Request.Model,
		CreatedAt: time.Now(),
		Response:  output,
		Done:      true,
	}
}
