# Signature using ens

Description
This is a program for verifying the ECDSA signature using the Ethereum blockchain. 
It gets a public key from the hash of the file and the signature. 
Next, it receives a hash from the public key or an address in Ethereum. 
Next, it accesses the Ethereum blockchain through the client 
and receives a domain (or name) by reverse resolving. 
Compares the entered name and the received one and outputs the result 
of the comparison

> 💡 You need to install [fyne](https://developer.fyne.io/started)
