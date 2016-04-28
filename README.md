# AMML
Array &amp; Map 's Markup Language, can be dead simple parsed into array and map(hashtable, dictionary).

## Features

- dead simple
- dead lightweight
- pure string
- useful everywhere

## Array

```
# This is an amml
1,word,4, a sentence
```

```csharp
// AMML -> "1,word,4, a sentence"
string[] arr = new string[]
{
  "1", "word", "4", "a sentence"
};
```

## Map
```
# This is an amml
k:123, j:word, t:a sentence
```

```python
# AMML -> "k:123, j:word, t:a sentence"
map_ = {"k":"123", "j":"word", "t":"sentence"}
```

### Array & Map Mix

```
# array inner map
k:123, j:(a,b,c,3)
```

```
# map inner array
1,word,(k:123,j:abc)
```
