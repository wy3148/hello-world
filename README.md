# hello-world example

## prerequisite
   docker 
   golang
## code
  git clone https://github.com/wy3148/hello-world.git
 
  cd hello-world
 
  go test -v
  
  docker build . -t hello-world
  
  docker run --rm -p 3000:3333 hello-world
  
