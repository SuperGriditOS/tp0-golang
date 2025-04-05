package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	log.Println("Hola soy un log")  // log. funciones de log

	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatal("Error al cargar la configuracion")
	}

	// loggeamos el valor de la config
	log.Printf("Configuracion cargada: %+v", globals.ClientConfig)

	//utils.LeerConsola() // leer de la consola el mensaje


	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto,globals.ClientConfig.Mensaje)

	// leer de la consola el mensaje
	utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
}
