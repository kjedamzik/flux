limit_fn = (table=<-) =>
  table
    |> range(start: 2018-05-22T19:00:00Z, stop: 2018-05-22T20:00:00Z)
    |> limit(n: 1)

testingTest(name: "limit", load: testLoadData, infile: "limit.in.csv", outfile: "limit.out.csv", test: limit_fn)
