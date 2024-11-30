import io
import matplotlib.pyplot as plt
import argparse
import sys

# Configuración del parser de argumentos
parser = argparse.ArgumentParser()

parser.add_argument('-t','--title',help='Titulo de la grafica')
parser.add_argument('-l','--labels',help='Etiquetas')
parser.add_argument('-d','--data',help='Datos')

args = parser.parse_args()

# Datos
# Convertir la entrada a una lista de números
valores = [float(valor.strip()) for valor in args.data.split(',')]

# Etiquetas para cada sector
etiquetas = [l for l in args.labels.split(',')]

# Crear la gráfica de pastel
plt.figure(figsize=(6, 6))
plt.pie(valores, labels=etiquetas, autopct='%1.1f%%', startangle=90)

# Añadir título si es proporcionado
if args.title:
    plt.title(args.title)

# Crear un buffer de bytes para guardar la imagen
buffer = io.BytesIO()

# Guardar la gráfica en el buffer en formato PNG
plt.savefig(buffer, format='png')

# Cerrar la figura para liberar recursos
plt.close()

# Mover el cursor al inicio del buffer
buffer.seek(0)

# Imprimir el contenido del buffer como PNG en stdout
sys.stdout.buffer.write(buffer.read())
