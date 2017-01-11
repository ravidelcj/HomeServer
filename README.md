# HomeServer
A Home Server which can host files on home wifi and unable other devices connected to the same wifi to download the file

First,
Add the files you want to serve to the publicFolder

Second,
run the Program using 

	go run ubuntuserver.go


Connect the device to the same network
Use your host computers Internal Ip as refernce

in the devices browser 
        
	open -> <ip address >:8008/getFile

This will list all the files in the publicFolder
click the file to download
