# CharDev OSINT Framework 🕵️‍♂️📞

Herramienta de inteligencia de fuentes abiertas (OSINT) enfocada en el análisis forense y la recolección de datos de números de teléfono. 

Inspirada en el concepto original de **PhoneInfoga**, pero reescrita desde cero con una arquitectura moderna (Go + Vue 3) para solucionar problemas de dependencias obsoletas y APIs caídas.

## 🚀 Módulos Integrados
Esta herramienta utiliza un enfoque BYOK (Bring Your Own Key) para garantizar la privacidad y estabilidad del usuario:
- **Veriphone:** Validación técnica, carrier y región geográfica.
- **IPQualityScore:** Análisis de riesgo, fraude y reputación de la línea.
- **Reporte IA (Llama 3):** Consolidación de datos en un informe forense en texto plano.

## ⚖️ Disclaimer Legal y Ético
**ESTA HERRAMIENTA ES ESTRICTAMENTE PARA FINES EDUCATIVOS Y DE INVESTIGACIÓN.** El creador no se hace responsable por el mal uso, daño o consecuencias legales derivadas de la utilización de este software. La recolección de información debe realizarse respetando las leyes locales e internacionales de privacidad. No utilices este software para acosar, vulnerar la privacidad o realizar actos ilegales. Al clonar o descargar este repositorio, asumes total responsabilidad por tus acciones.

## 🛠️ Instalación y Uso (Guía para Colaboradores)
El proyecto está dividido en dos partes: el servidor backend en Go y la interfaz frontend en Vue 3 (Vite). Para correr la herramienta en tu entorno local, sigue estos pasos en tu terminal:

**Requisitos Previos**
- Tener instalado Go (Golang) v1.18 o superior.
- Tener instalado Node.js y npm.

**1. Levantar el Frontend (Interfaz de Usuario)**
Abre una terminal en la raíz del proyecto y ejecuta:
```bash
cd web
npm install
npm run dev
```
Esto levantará el servidor de desarrollo de Vite (generalmente en `http://localhost:5173`).

**2. Levantar el Proxy Forense (Backend Go)**
Abre una nueva pestaña en tu terminal (sin cerrar la del frontend), ve a la raíz del proyecto y ejecuta:
```bash
go run main.go --web
```
Esto iniciará el servidor backend en `http://localhost:5000`, el cual se comunicará con las APIs de OSINT.

**3. Configuración BYOK (Bring Your Own Key)**
- Abre tu navegador en la dirección que te dio Vite (`http://localhost:5173`).
- Haz clic en '⚙️ BYOK Config'.
- Ingresa tus llaves gratuitas de Veriphone e IPQualityScore, y tu llave de Groq para el reporte de IA.

¡Listo para investigar!

## 🤝 Colaboración y Mantenimiento
Este framework nació como una investigación personal de ciberseguridad. Actualmente, debido a mis estudios universitarios, no dispongo del tiempo necesario para mantener el proyecto de forma activa a largo plazo.

El código es totalmente libre. Cualquier desarrollador o investigador de OSINT está invitado a hacer **Fork**, proponer **Pull Requests**, sumar nuevas APIs o tomar la posta del mantenimiento. Siéntanse libres de romper, mejorar y expandir este código.

---
**Powered by Char Dev Quantum**
