package http

type HTTPTestVariantResult struct {
	Success bool
	Err     error
	Log     []byte
	KV      map[string]string
}

type HTTPTestVariant struct {
	Name      string
	Arguments map[string]any
	Result    HTTPTestVariantResult
	Expected  string
}

type HTTPTest struct {
	Name     string
	Cmd      string
	Variants []HTTPTestVariant
}

func (ht HTTPTest) Run() error    { return nil }
func (ht HTTPTest) Success() bool { return true }

func AllHTTPTests() []HTTPTest {
	return []HTTPTest{
		{
			Name: "Resolve A record",
			Cmd:  "/usr/bin/dig",
			Variants: []HTTPTestVariant{
				{
					Name: "... of something.example.com",
					Arguments: map[string]any{
						"args": []string{"A", "something.example.com"},
					},
				},
				{
					Name: "... of xyz.example.com",
					Arguments: map[string]any{
						"args": []string{"A", "xyz.example.com"},
					},
				},
				{
					Name: "... of abc.example.com",
					Arguments: map[string]any{
						"args": []string{"A", "abc.example.com"},
					},
				},
			},
		},
		{
			Name: "Resolve AAAA record",
			Cmd:  "/usr/bin/dig",
			Variants: []HTTPTestVariant{
				{
					Name: "... of something.example.com",
					Arguments: map[string]any{
						"args": []string{"AAAA", "something.example.com"},
					},
				},
				{
					Name: "... of xyz.example.com",
					Arguments: map[string]any{
						"args": []string{"AAAA", "xyz.example.com"},
					},
				},
				{
					Name: "... of abc.example.com",
					Arguments: map[string]any{
						"args": []string{"AAAA", "abc.example.com"},
					},
				},
			},
		},
		{
			Name: "Resolve MX record",
			Cmd:  "/usr/bin/dig",
			Variants: []HTTPTestVariant{
				{
					Name: "... of something.example.com",
					Arguments: map[string]any{
						"args": []string{"MX", "something.example.com"},
					},
				},
				{
					Name: "... of xyz.example.com",
					Arguments: map[string]any{
						"args": []string{"MX", "xyz.example.com"},
					},
				},
				{
					Name: "... of abc.example.com",
					Arguments: map[string]any{
						"args": []string{"MX", "abc.example.com"},
					},
				},
			},
		},
	}
}
