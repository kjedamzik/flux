// DO NOT EDIT: This file is autogenerated via the builtin command.

package telegraf

import (
	flux "github.com/influxdata/flux"
	ast "github.com/influxdata/flux/ast"
)

func init() {
	flux.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 48,
					Line:   8,
				},
				File:   "system.flux",
				Source: "package telegraf\n\nimport \"influxdb\"\n\noption bucket = \"telegraf\"\n\ncpu = (bucket=bucket) =>\n    influxdb.measurement(bucket:bucket,m:\"cpu\")",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.OptionStatement{
			Assignment: &ast.VariableAssignment{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 27,
							Line:   5,
						},
						File:   "system.flux",
						Source: "bucket = \"telegraf\"",
						Start: ast.Position{
							Column: 8,
							Line:   5,
						},
					},
				},
				ID: &ast.Identifier{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 14,
								Line:   5,
							},
							File:   "system.flux",
							Source: "bucket",
							Start: ast.Position{
								Column: 8,
								Line:   5,
							},
						},
					},
					Name: "bucket",
				},
				Init: &ast.StringLiteral{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 27,
								Line:   5,
							},
							File:   "system.flux",
							Source: "\"telegraf\"",
							Start: ast.Position{
								Column: 17,
								Line:   5,
							},
						},
					},
					Value: "telegraf",
				},
			},
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 27,
						Line:   5,
					},
					File:   "system.flux",
					Source: "option bucket = \"telegraf\"",
					Start: ast.Position{
						Column: 1,
						Line:   5,
					},
				},
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 48,
						Line:   8,
					},
					File:   "system.flux",
					Source: "cpu = (bucket=bucket) =>\n    influxdb.measurement(bucket:bucket,m:\"cpu\")",
					Start: ast.Position{
						Column: 1,
						Line:   7,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 4,
							Line:   7,
						},
						File:   "system.flux",
						Source: "cpu",
						Start: ast.Position{
							Column: 1,
							Line:   7,
						},
					},
				},
				Name: "cpu",
			},
			Init: &ast.FunctionExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 48,
							Line:   8,
						},
						File:   "system.flux",
						Source: "(bucket=bucket) =>\n    influxdb.measurement(bucket:bucket,m:\"cpu\")",
						Start: ast.Position{
							Column: 7,
							Line:   7,
						},
					},
				},
				Body: &ast.CallExpression{
					Arguments: []ast.Expression{&ast.ObjectExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 47,
									Line:   8,
								},
								File:   "system.flux",
								Source: "bucket:bucket,m:\"cpu\"",
								Start: ast.Position{
									Column: 26,
									Line:   8,
								},
							},
						},
						Properties: []*ast.Property{&ast.Property{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 39,
										Line:   8,
									},
									File:   "system.flux",
									Source: "bucket:bucket",
									Start: ast.Position{
										Column: 26,
										Line:   8,
									},
								},
							},
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 32,
											Line:   8,
										},
										File:   "system.flux",
										Source: "bucket",
										Start: ast.Position{
											Column: 26,
											Line:   8,
										},
									},
								},
								Name: "bucket",
							},
							Value: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 39,
											Line:   8,
										},
										File:   "system.flux",
										Source: "bucket",
										Start: ast.Position{
											Column: 33,
											Line:   8,
										},
									},
								},
								Name: "bucket",
							},
						}, &ast.Property{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 47,
										Line:   8,
									},
									File:   "system.flux",
									Source: "m:\"cpu\"",
									Start: ast.Position{
										Column: 40,
										Line:   8,
									},
								},
							},
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 41,
											Line:   8,
										},
										File:   "system.flux",
										Source: "m",
										Start: ast.Position{
											Column: 40,
											Line:   8,
										},
									},
								},
								Name: "m",
							},
							Value: &ast.StringLiteral{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 47,
											Line:   8,
										},
										File:   "system.flux",
										Source: "\"cpu\"",
										Start: ast.Position{
											Column: 42,
											Line:   8,
										},
									},
								},
								Value: "cpu",
							},
						}},
					}},
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 48,
								Line:   8,
							},
							File:   "system.flux",
							Source: "influxdb.measurement(bucket:bucket,m:\"cpu\")",
							Start: ast.Position{
								Column: 5,
								Line:   8,
							},
						},
					},
					Callee: &ast.MemberExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 25,
									Line:   8,
								},
								File:   "system.flux",
								Source: "influxdb.measurement",
								Start: ast.Position{
									Column: 5,
									Line:   8,
								},
							},
						},
						Object: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 13,
										Line:   8,
									},
									File:   "system.flux",
									Source: "influxdb",
									Start: ast.Position{
										Column: 5,
										Line:   8,
									},
								},
							},
							Name: "influxdb",
						},
						Property: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 25,
										Line:   8,
									},
									File:   "system.flux",
									Source: "measurement",
									Start: ast.Position{
										Column: 14,
										Line:   8,
									},
								},
							},
							Name: "measurement",
						},
					},
				},
				Params: []*ast.Property{&ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 21,
								Line:   7,
							},
							File:   "system.flux",
							Source: "bucket=bucket",
							Start: ast.Position{
								Column: 8,
								Line:   7,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 14,
									Line:   7,
								},
								File:   "system.flux",
								Source: "bucket",
								Start: ast.Position{
									Column: 8,
									Line:   7,
								},
							},
						},
						Name: "bucket",
					},
					Value: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 21,
									Line:   7,
								},
								File:   "system.flux",
								Source: "bucket",
								Start: ast.Position{
									Column: 15,
									Line:   7,
								},
							},
						},
						Name: "bucket",
					},
				}},
			},
		}},
		Imports: []*ast.ImportDeclaration{&ast.ImportDeclaration{
			As: nil,
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 18,
						Line:   3,
					},
					File:   "system.flux",
					Source: "import \"influxdb\"",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			Path: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 18,
							Line:   3,
						},
						File:   "system.flux",
						Source: "\"influxdb\"",
						Start: ast.Position{
							Column: 8,
							Line:   3,
						},
					},
				},
				Value: "influxdb",
			},
		}},
		Name: "system.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 17,
						Line:   1,
					},
					File:   "system.flux",
					Source: "package telegraf",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 17,
							Line:   1,
						},
						File:   "system.flux",
						Source: "telegraf",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "telegraf",
			},
		},
	}},
	Package: "telegraf",
	Path:    "influxdb/telegraf",
}
