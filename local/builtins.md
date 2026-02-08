# I/O

```js
input()
readLine()
readFile(path)
writeFile(path, data)
appendFile(path, data)
exists(path)
deleteFile(path)
```

# Matemática

```js
abs(x)
pow(a, b)
sqrt(x)
cbrt(x)
floor(x)
ceil(x)
round(x)
min(a, b)
max(a, b)
clamp(x, min, max)
rand()
randInt(min, max)
sin(x)
cos(x)
tan(x)
log(x)
log10(x)
exp(x)
```

# Strings

```js
len(s)
upper(s)
lower(s)
trim(s)
split(s, sep)
join(arr, sep)
replace(s, old, new)
contains(s, sub)
startsWith(s, sub)
endsWith(s, sub)
indexOf(s, sub)
lastIndexOf(s, sub)
substring(s, start, end)
charAt(s, i)
repeat(s, n)
```

# Arrays
```js
push(arr, v)
pop(arr)
shift(arr)
unshift(arr, v)
insert(arr, i, v)
remove(arr, i)
slice(arr, start, end)
concat(a, b)
reverse(arr)
sort(arr)
map(arr, fn)
filter(arr, fn)
reduce(arr, fn, init)
find(arr, fn)
index(arr, v)
contains(arr, v)
unique(arr)
shuffle(arr)
```

# Objetos / Mapas

```js
keys(obj)
values(obj)
entries(obj)
hasKey(obj, key)
deleteKey(obj, key)
merge(a, b)
clone(obj)
size(obj)
clear(obj)
```

# Tipos / Conversão
```js
type(x)
isNumber(x)
isString(x)
isBool(x)
isArray(x)
isObject(x)
isNil(x)
toInt(x)
toFloat(x)
toString(x)
toBool(x)
```

# Tempo / Sistema
```js
now()
sleep(ms)
clock()
env(name)
exit(code)
args()
```

# Funcionais
```js
call(fn, args)
apply(fn, arr)
partial(fn, args)
once(fn)
memo(fn)
```

# Hash / Encode
```js
hash(s)
sha1(s)
sha256(s)
base64Encode(s)
base64Decode(s)
hexEncode(s)
hexDecode(s)
```

# Erros / Debug
```js
panic(msg)
assert(cond, msg)
trace()
dump(x)
inspect(x)
```

# JSON / Serialização
```js
jsonParse(s)
jsonStringify(x)
serialize(x)
deserialize(s)
```

# Regex
```js
match(s, pattern)
findAll(s, pattern)
replaceRegex(s, pattern, repl)
splitRegex(s, pattern)
```

# Iteração
```js
range(n)
range(a, b)
range(a, b, step)
enumerate(arr)
zip(a, b)
```

# Runtime
```js
gc()
memUsage()
version()
platform()
```
