import socket

HOST = '127.0.0.1'      # Symbolic name meaning the local host
PORT = 1235            # Arbitrary non-privileged port

socketRegister = socket.socket()
socketRegister.connect((HOST,1236))
socketRegister.send("action=register id=pythonStub port=1235")


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind((HOST, PORT))
s.listen(1)

while 1:
	conn, addr = s.accept()
	print 'Connected by', conn
        data = conn.recv(1024)
        if not data:
                break
        conn.send("id=mainHeater")

conn.close()
