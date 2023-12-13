package person

import "time"

type Person struct {
	Id         int        `json:"id"`
	Apelido    string     `json:"apelido"`
	Nome       string     `json:"nome"`
	Nascimento customDate `json:"nascimento"`
	Stack      []string   `json:"stack"`
}

type customDate struct {
	time.Time
}

func (t *customDate) UnmarshalJSON(b []byte) error {
	time, err := time.Parse(`"2006-01-02"`, string(b))
	t.Time = time

	return err
}

func (t *customDate) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(`"2006-01-02"`)), nil
}
