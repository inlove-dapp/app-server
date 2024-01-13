package clients

import (
	"errors"
	"testing"
)

func TestParseError(t *testing.T) {
	client := &PrismaErrorClient{}
	testCases := []struct {
		name        string
		err         string
		wantErrCode string
		wantKind    string
	}{
		{
			name:        "Test Case 1",
			err:         `pql error: Error occurred during query execution: ConnectorError(ConnectorError { user_facing_error: Some(KnownError { message: "Unique constraint failed on the constraint: ` + "`User_supabaseUserId_key`" + `", meta: Object {"target": String("User_supabaseUserId_key")}, error_code: "P2002" }), kind: UniqueConstraintViolation { constraint: Index("User_supabaseUserId_key") }, transient: false })`,
			wantErrCode: "P2002",
			wantKind:    "UniqueConstraintViolation",
		},
		{
			name:        "Test Case 2",
			err:         `pql error: Error occurred during query execution: ConnectorError(ConnectorError { user_facing_error: Some(KnownError { message: "Unique constraint failed on the constraint: ` + "`AnotherConstraint_key`" + `", meta: Object {"target": String("AnotherConstraint_key")}, error_code: "P2003" }), kind: ForeignKeyViolation { constraint: Index("AnotherConstraint_key") }, transient: false })`,
			wantErrCode: "P2003",
			wantKind:    "ForeignKeyViolation",
		},
		{
			name:        "Test Case 3",
			err:         `pql error: Error occurred during query execution: ConnectorError(ConnectorError { user_facing_error: Some(KnownError { message: "Unique constraint failed on the constraint: ` + "`YetAnotherConstraint_key`" + `", meta: Object {"target": String("YetAnotherConstraint_key")}, error_code: "P2004" }), kind: NullConstraintViolation { constraint: Index("YetAnotherConstraint_key") }, transient: false })`,
			wantErrCode: "P2004",
			wantKind:    "NullConstraintViolation",
		},
		{
			name:        "Test Case 3",
			err:         `pql error:`,
			wantErrCode: "",
			wantKind:    "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotErrCode, gotKind := client.parseError(tc.err)
			if gotErrCode != tc.wantErrCode || gotKind != tc.wantKind {
				t.Errorf("parseError(%q) = %q, %q; want %q, %q", tc.err, gotErrCode, gotKind, tc.wantErrCode, tc.wantKind)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	client := NewPrismaErrorClient()

	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "Record already exists",
			err:      errors.New(`error_code: "P2002"`),
			expected: "Record already exists",
		},
		{
			name:     "Foreign key constraint failed",
			err:      errors.New(`error_code: "P2003"`),
			expected: "Foreign key constraint failed",
		},
		{
			name:     "Null constraint failed",
			err:      errors.New(`error_code: "P2004"`),
			expected: "Null constraint failed",
		},
		{
			name:     "Unknown error code",
			err:      errors.New(`pql error: Error occurred during query execution: ConnectorError(ConnectorError { user_facing_error: Some(KnownError { message: "Unique constraint failed on the constraint: ` + "`YetAnotherConstraint_key`" + `", meta: Object {"target": String("YetAnotherConstraint_key")}, error_code: "P9999" }), kind: Unknown { constraint: Index("YetAnotherConstraint_key") }, transient: false })`),
			expected: "P9999 Unknown",
		},
		{
			name:     "Unknown error",
			err:      errors.New(`pql error`),
			expected: "pql error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := client.HandleError(tt.err); got != tt.expected {
				t.Errorf("HandleError() = %v, want %v", got, tt.expected)
			}
		})
	}
}
