// DO NOT EDIT: This file is autogenerated via the builtin command.

package testing

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
					Column: 58,
					Line:   4,
				},
				Source: "package testing\n\ntest = (name,input,want,test) =>\n    assertEquals(name:name, got:input|>test(), want:want)",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 58,
						Line:   4,
					},
					Source: "test = (name,input,want,test) =>\n    assertEquals(name:name, got:input|>test(), want:want)",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 5,
							Line:   3,
						},
						Source: "test",
						Start: ast.Position{
							Column: 1,
							Line:   3,
						},
					},
				},
				Name: "test",
			},
			Init: &ast.FunctionExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 58,
							Line:   4,
						},
						Source: "(name,input,want,test) =>\n    assertEquals(name:name, got:input|>test(), want:want)",
						Start: ast.Position{
							Column: 8,
							Line:   3,
						},
					},
				},
				Body: &ast.CallExpression{
					Arguments: []ast.Expression{&ast.ObjectExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 57,
									Line:   4,
								},
								Source: "name:name, got:input|>test(), want:want",
								Start: ast.Position{
									Column: 18,
									Line:   4,
								},
							},
						},
						Properties: []*ast.Property{&ast.Property{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 27,
										Line:   4,
									},
									Source: "name:name",
									Start: ast.Position{
										Column: 18,
										Line:   4,
									},
								},
							},
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 22,
											Line:   4,
										},
										Source: "name",
										Start: ast.Position{
											Column: 18,
											Line:   4,
										},
									},
								},
								Name: "name",
							},
							Value: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 27,
											Line:   4,
										},
										Source: "name",
										Start: ast.Position{
											Column: 23,
											Line:   4,
										},
									},
								},
								Name: "name",
							},
						}, &ast.Property{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 46,
										Line:   4,
									},
									Source: "got:input|>test()",
									Start: ast.Position{
										Column: 29,
										Line:   4,
									},
								},
							},
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 32,
											Line:   4,
										},
										Source: "got",
										Start: ast.Position{
											Column: 29,
											Line:   4,
										},
									},
								},
								Name: "got",
							},
							Value: &ast.PipeExpression{
								Argument: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 38,
												Line:   4,
											},
											Source: "input",
											Start: ast.Position{
												Column: 33,
												Line:   4,
											},
										},
									},
									Name: "input",
								},
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 46,
											Line:   4,
										},
										Source: "input|>test()",
										Start: ast.Position{
											Column: 33,
											Line:   4,
										},
									},
								},
								Call: &ast.CallExpression{
									Arguments: nil,
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 46,
												Line:   4,
											},
											Source: "test()",
											Start: ast.Position{
												Column: 40,
												Line:   4,
											},
										},
									},
									Callee: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 44,
													Line:   4,
												},
												Source: "test",
												Start: ast.Position{
													Column: 40,
													Line:   4,
												},
											},
										},
										Name: "test",
									},
								},
							},
						}, &ast.Property{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 57,
										Line:   4,
									},
									Source: "want:want",
									Start: ast.Position{
										Column: 48,
										Line:   4,
									},
								},
							},
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 52,
											Line:   4,
										},
										Source: "want",
										Start: ast.Position{
											Column: 48,
											Line:   4,
										},
									},
								},
								Name: "want",
							},
							Value: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 57,
											Line:   4,
										},
										Source: "want",
										Start: ast.Position{
											Column: 53,
											Line:   4,
										},
									},
								},
								Name: "want",
							},
						}},
					}},
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 58,
								Line:   4,
							},
							Source: "assertEquals(name:name, got:input|>test(), want:want)",
							Start: ast.Position{
								Column: 5,
								Line:   4,
							},
						},
					},
					Callee: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 17,
									Line:   4,
								},
								Source: "assertEquals",
								Start: ast.Position{
									Column: 5,
									Line:   4,
								},
							},
						},
						Name: "assertEquals",
					},
				},
				Params: []*ast.Property{&ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 13,
								Line:   3,
							},
							Source: "name",
							Start: ast.Position{
								Column: 9,
								Line:   3,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 13,
									Line:   3,
								},
								Source: "name",
								Start: ast.Position{
									Column: 9,
									Line:   3,
								},
							},
						},
						Name: "name",
					},
					Value: nil,
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 19,
								Line:   3,
							},
							Source: "input",
							Start: ast.Position{
								Column: 14,
								Line:   3,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 19,
									Line:   3,
								},
								Source: "input",
								Start: ast.Position{
									Column: 14,
									Line:   3,
								},
							},
						},
						Name: "input",
					},
					Value: nil,
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 24,
								Line:   3,
							},
							Source: "want",
							Start: ast.Position{
								Column: 20,
								Line:   3,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 24,
									Line:   3,
								},
								Source: "want",
								Start: ast.Position{
									Column: 20,
									Line:   3,
								},
							},
						},
						Name: "want",
					},
					Value: nil,
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 29,
								Line:   3,
							},
							Source: "test",
							Start: ast.Position{
								Column: 25,
								Line:   3,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 29,
									Line:   3,
								},
								Source: "test",
								Start: ast.Position{
									Column: 25,
									Line:   3,
								},
							},
						},
						Name: "test",
					},
					Value: nil,
				}},
			},
		}},
		Imports: nil,
		Name:    "test.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 16,
						Line:   1,
					},
					Source: "package testing",
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
							Column: 16,
							Line:   1,
						},
						Source: "testing",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "testing",
			},
		},
	}},
	Package: "testing",
	Path:    "testing",
}