# additionalProperties

See https://github.com/swaggo/swag/issues/374

The additionalProperties work fine when I use it in  `@Success`

See https://github.com/dvasilen/additionalProperties/blob/master/src/test/main.go#L26

But When I try to use them in the  `@Param` for the query the parsing fails

```
// @Param   q query {object} model.PropMap false "The optional query parameters"

```

```
2019/04/12 13:56:31 ParseComment error in file test/main.go :missing required param comment parameters "q query {object} model.PropMap false "The optional query parameters""
```

