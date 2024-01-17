# Climer

<div>
	<img width="40" src="https://raw.githubusercontent.com/tandpfun/skill-icons/de91fca307a83d75fc5b1f6ce24540454acead41/icons/Arduino.svg" />
	<img width="40" src="https://raw.githubusercontent.com/tandpfun/skill-icons/de91fca307a83d75fc5b1f6ce24540454acead41/icons/GoLang.svg" />
	<img width="40" src="https://raw.githubusercontent.com/tandpfun/skill-icons/de91fca307a83d75fc5b1f6ce24540454acead41/icons/JavaScript.svg" />
</div>

## O Projeto

Uma aplicação que utiliza o sensor DHT11 para medir dados de __Temperatura__ e __Umidade__ do ambiente.  

Esses dados são enviados para o __Arduino Uno__ que através do __módulo Ethernet ENC28J60__ monta uma requisição http com os dados e envia pra rede.

Criei um servidor com a linguagem __Golang__ 

---
### Tecnologias Utilizadas

| Tecnologias | Funções |
| --- | --- |
| IOT ||
| <img width="100" alt="Módulo DHT11" src="https://cdn.awsli.com.br/800x800/78/78150/produto/3440956/modulo_umidade_temperatura_dht11-z97e6uxojd.png" style="margin:30px;"/> |  _* DHT11_<br/>Sensor de __temperatura__ e __umidade__ local |
| <img width="100" alt="Módulo Ethernet ENC28j60" src="https://images.tcdn.com.br/img/img_prod/557243/modulo_ethernet_enc28j60_mini_560_1_a82d51d7e4ce1d966f0af28cb7e36043.png" style="margin:30px;"/> | _* Módulo ENC28j60_<br/>Módulo de conexão Ethernet, envia a requisição http para a rede |
| <img src="https://upload.wikimedia.org/wikipedia/commons/1/1c/Arduino-uno.jpg" width="100" alt="Arduino Uno" style="margin:30px;"/> | _* Arduino Uno_<br/>Microcontrolador que recebe os dados  |