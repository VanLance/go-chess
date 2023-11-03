const testApi = async () => {
  const req = await fetch("http://localhost:8080/")
  const data = await req.json()
  console.log(data)
}

(async ()=>{ await testApi()})()