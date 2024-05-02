Secuencial = []
Concurrente = []

with open('Resultados/1. SecuencialResultado.txt', 'r') as f:
    for linea in f:
        valor = int(linea.strip())  
        Secuencial.append(valor)

valores_ordenados = sorted(Secuencial)

valores_recortados = valores_ordenados[50:-50]
media_recortada = sum(valores_recortados) / len(valores_recortados)

print("Media recortada Secuencial:", media_recortada)


with open('Resultados/2. ConcurrenteResultado.txt', 'r') as f:
    for linea in f:
        valor = int(linea.strip())  
        Concurrente.append(valor)

valores_ordenados = sorted(Concurrente)

valores_recortados = valores_ordenados[50:-50]
media_recortada = sum(valores_recortados) / len(valores_recortados)

print("Media recortada Concurrente:", media_recortada)
