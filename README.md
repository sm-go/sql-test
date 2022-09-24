# sql-test

## single row query

## emp project structrue

```
** folder တစ်ခုအတွင်းထဲမှာ package တစ်ခုပဲ ရှိရမယ်။ တူလို့ မရပါ။ package တစ်ခုကိုမှ file ကြိုက်သလောက်ခွဲလို့ရ၏ (folder ထဲမှာ package တူ, file ကွဲ လို့ ရ၏)
eg. 
 - src (folder)
   |__ main.go (package main)
   |__another.go (package main)

ဒါမျိုး structure ဆိုရင် $ go run folder/*.go နဲ့ run သင့်။
သို့သော် project structure မကျပါ။
Project structure ကျ ဖို့ ဆိုရင် main.go က သီးသန့် ဖြစ်သင့်
$ go run main.go လို့ ဖြစ်သင့်

ဒါပေမဲ့ ဒီ project ထဲမှာ package ခွဲဖို့ folder မခွဲပဲ တစ် ဖိုဒါထဲမှာပဲ main package အောက် အကုန် ဝင် run 

```