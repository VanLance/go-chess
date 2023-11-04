let pieces


const testApi = async () => {
  const res = await fetch("http://localhost:8080/")

  const data = await res.json()
  pieces = data
  return data
}

(async () => { console.log(await testApi()) })();
(async () => { (await firstMove()) })()


async function firstMove() {
  const pieces = await testApi()
  const res = await fetch("http://localhost:8080/make-move", {
    method: "POST",
    headers: {
      "Content-Type": 'application/json'
    },
    body: JSON.stringify({
      previousState: [...pieces.PlayerOnePieces, ...pieces.PlayerTwoPieces],
      move: {
        startingPosition: "A2",
        landingPosition: "A3",
        player: 1
      }
    })
  })
  // if (res.ok) {
  //   const data = await res.json();
  //   console.log("Response Data:", data);
  // } else {
  //   console.error('HTTP error:', res.status);
  // }
}

