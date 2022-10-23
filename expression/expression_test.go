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
			in: "out=ADD 2",
			result: TestTranslateResult{
				out: "",
				err: ArgsValidationError,
			},
		},
		{
			in: "out=SET FOO|>ADD 2",
			result: TestTranslateResult{
				out: "out=((FOO)+2)",
				err: nil,
			},
		},
		{
			in: "out=SUB FOO 2",
			result: TestTranslateResult{
				out: "out=(FOO-2)",
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
				err: OperationNotFound,
			},
		},
		{
			in: "BULSHIT",
			result: TestTranslateResult{
				out: "",
				err: ExpressionDividerNotFound,
			},
		},
		{
			in: "BULSHIT=",
			result: TestTranslateResult{
				out: "",
				err: OperationNotFound,
			},
		},
		{
			in: "out=ADD|>SUB 1",
			result: TestTranslateResult{
				out: "",
				err: ArgsValidationError,
			},
		},
	}
	for i, tc := range tests {
		res, err := Translate(tc.in)
		if err != tc.result.err || res != tc.result.out {
			t.Errorf("[%d] [error] in=\"%s\", \"%s\"=\"%s\", %v=%v", i, tc.in, tc.result.out, res, err, tc.result.err)
			continue
		}
		t.Logf("[%d] [success] in=\"%s\", our=\"%s\", err=%v", i, tc.in, tc.result.out, tc.result.err)
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
