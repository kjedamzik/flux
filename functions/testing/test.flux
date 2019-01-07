package testing

test = (name,input,want,test) =>
    assertEquals(name: name, got: input|>test(), want: want)
