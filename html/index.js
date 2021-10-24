import * as nats from './nats.js'

document.addEventListener('DOMContentLoaded', setupListeners)

async function setupListeners() {
  const nc = await nats.connect({ servers: 'ws://localhost:9222' })
  const sc = nats.StringCodec()
  document.getElementById('stormdeck').addEventListener('click', () => {
    nc.request('GAME.0.DRAW.storm', '')
      .then(msg => { console.log(sc.decode(msg.data)) })
  })
}

