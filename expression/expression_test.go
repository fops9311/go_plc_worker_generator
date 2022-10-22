package expression

import "testing"

func TestTranslate(t *testing.T) {
	tests := []TestTranslateCase{
		{
			in: "out=ADD 1 2",
			result: TestTranslateResult{
				out: "out=(1+2)",
				err: nil,
			},
		},
		{
			in: "out=SUB 2",
			result: TestTranslateResult{
				out: "out=(-2)",
				err: nil,
			},
		},
		{
			in: "out=MUL 1 2",
			result: TestTranslateResult{
				out: "out=(1*2)",
				err: nil,
			},
		},
		{
			in: "out=DIV 1 2",
			result: TestTranslateResult{
				out: "out=(1/2)",
				err: nil,
			},
		},
		{
			in: "out=ADD 1 2|>SUB 1 2|>MUL 1 2|>DIV 1 2",
			result: TestTranslateResult{
				out: "out=((((1+2)-1-2)*1*2)/1/2)",
				err: nil,
			},
		},
		{
			in: "out=BULSHIT 1",
			result: TestTranslateResult{
				out: "",
				err: ExpressionError,
			},
		},
		{
			in: "BULSHIT",
			result: TestTranslateResult{
				out: "",
				err: ExpressionError,
			},
		},
		{
			in: "BULSHIT=",
			result: TestTranslateResult{
				out: "",
				err: ExpressionError,
			},
		},
		{
			in: "out=BULSHIT",
			result: TestTranslateResult{
				out: "",
				err: ExpressionError,
			},
		},
		{
			in: "out=ADD|>SUB 1",
			result: TestTranslateResult{
				out: "out=()",
				err: nil,
			},
		},
	}
	for _, tc := range tests {
		res, err := Translate(tc.in)
		if err != tc.result.err || res != tc.result.out {
			t.Errorf("[error] in=\"%s\", \"%s\"=\"%s\", err=%v", tc.in, tc.result.out, res, tc.result.err)
			continue
		}
		t.Logf("[success] in=\"%s\", our=\"%s\", err=%v", tc.in, tc.result.out, tc.result.err)
	}

}

type TestTranslateCase struct {
	in     string
	result TestTranslateResult
}
type TestTranslateResult struct {
	out string
	err error
}
