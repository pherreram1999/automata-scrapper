import io
import matplotlib.pyplot as plt
import argparse
import sys

# Configuración del parser de argumentos
parser = argparse.ArgumentParser()

parser.add_argument('-t', '--title', help='Titulo de la grafica')
parser.add_argument('-l', '--labels', help='Etiquetas')
parser.add_argument('-d', '--data', help='Datos')

args = parser.parse_args()

# Datos
# Convertir la entrada a una lista de números
valores = [float(valor.strip()) for valor in args.data.split(',')]

# Etiquetas para cada sector
etiquetas = [l for l in args.labels.split(',')]

# Definir colores personalizados (uno para cada etiqueta)
colores = plt.cm.Paired(range(len(etiquetas)))  # Usando una paleta de colores

# Crear la gráfica de pastel
plt.figure(figsize=(8, 8))  # Aumentamos el tamaño de la figura para que haya más espacio

# Creamos el gráfico de pastel y pasamos los colores
wedges, texts, autotexts = plt.pie(valores, labels=etiquetas, autopct='%1.1f%%', startangle=90, colors=colores)

# Añadir título si es proporcionado
if args.title:
    plt.title(args.title)

# Añadir la leyenda al lado con colores correspondientes
plt.legend(wedges, etiquetas, title="Etiquetas", loc="center left", bbox_to_anchor=(1, 0, 0.5, 1))

# Crear un buffer de bytes para guardar la imagen
buffer = io.BytesIO()

# Guardar la gráfica en el buffer en formato PNG, con la opción bbox_inches='tight' para evitar recortes
plt.savefig(buffer, format='png', bbox_inches='tight')

# Cerrar la figura para liberar recursos
plt.close()

# Mover el cursor al inicio del buffer
buffer.seek(0)

# Imprimir el contenido del buffer como PNG en stdout
sys.stdout.buffer.write(buffer.read())
