package mist_error

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func LogError(ctx context.Context, e diag.Diagnostics, message string) {
	for _, et := range e.Errors() {
		var ei interface{} = et
		ef := ei.(diag.Diagnostic)
		tflog.Error(ctx, message, map[string]interface{}{
			"Detail":  ef.Detail(),
			"Summary": ef.Summary(),
		})
	}
}
