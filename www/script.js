const log_status = document.getElementById("log_status")
const graph_temperatura = document.getElementById("graph-temperatura")
const graph_umidade = document.getElementById("graph-umidade")
const log_temperatura = document.getElementById("log-temperatura")
const log_umidade = document.getElementById("log-umidade")

const graph_max = 300
const max_bar_num = 6

var socket = new WebSocket(`ws://${location.host}/echo`)

const historyTemperatura = []
const historyUmidade = []

socket.onopen = function(){
	log_status.textContent = "Conectado"
	log_status.style.color = "green"

	socket.send("echo")
}

socket.onmessage = function({data}){
	console.log( data )

	if ( data == "echo" ) {
		log_status.innerText = "Conectado"
		log_status.style.color = "#32de84"
		return
	}

	const cookedData = JSON.parse( data )
	drawGraph(cookedData)
}

socket.onclose = function(){
	log_status.textContent = "Conexão Finalizada"
	log_status.style.color = "red"
}

function drawGraph({ temperatura, umidade }) {
	historyTemperatura.push( temperatura )
	if ( historyTemperatura.length > max_bar_num )
		historyTemperatura.shift()

	historyUmidade.push( umidade )
	if ( historyUmidade.length > max_bar_num )
		historyUmidade.shift()

	graph_temperatura.innerHTML = ""
	historyTemperatura.forEach(d => {
		const graph_bar_1 = document.createElement("div")
		graph_bar_1.className = "graph-bar"
		graph_bar_1.style.width = `${d/50 * graph_max}px`
		graph_bar_1.innerText = d
		graph_temperatura.appendChild( graph_bar_1 )
	})

	graph_umidade.innerHTML = ""
	historyUmidade.forEach(d => {
		const graph_bar_1 = document.createElement("div")
		graph_bar_1.className = "graph-bar"
		graph_bar_1.style.width = `${d/100 * graph_max}px`
		graph_bar_1.innerText = d
		graph_umidade.appendChild( graph_bar_1 )
	})

	log_temperatura.textContent = temperatura
	log_umidade.textContent = umidade
}