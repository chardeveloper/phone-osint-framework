<!-- @author CHAR DEV QUANTUM -->
<template>
  <div class="osint-container min-vh-100 p-4 bg-dark text-light">
    <header class="d-flex justify-content-between align-items-center mb-5 border-bottom border-secondary pb-3">
      <h1 class="text-primary m-0">CharDev OSINT Framework</h1>
      <button class="btn btn-outline-info" data-bs-toggle="modal" data-bs-target="#configModal">⚙️ BYOK Config</button>
    </header>

    <main class="container">
      <div class="row justify-content-center mb-5">
        <div class="col-md-8">
          <div class="input-group input-group-lg">
            <input v-model="targetPhone" type="text" class="form-control bg-secondary text-white border-dark" placeholder="Ingresa el número (ej: +549...)">
          </div>
        </div>
      </div>

      <div class="row g-4 justify-content-center mb-5">
        <!-- Veriphone -->
        <div class="col-md-5">
          <div class="card h-100 bg-dark border-secondary shadow-sm">
            <div class="card-header border-secondary d-flex justify-content-between">
              <span class="fw-bold">Veriphone</span>
              <span :class="veriphoneKey ? 'text-success' : 'text-danger'">{{ veriphoneKey ? 'Ready' : 'Missing Key' }}</span>
            </div>
            <div class="card-body d-flex flex-column">
              <button class="btn btn-sm btn-primary w-100 mb-3" :disabled="!veriphoneKey || !targetPhone || loadingVeriphone" @click="ejecutarVeriphone">
                <span v-if="loadingVeriphone" class="spinner-border spinner-border-sm me-2"></span>
                Execute Veriphone
              </button>
              <div v-if="resultadoVeriphone" class="flex-grow-1 bg-secondary rounded p-2 overflow-auto" style="max-height: 250px;">
                <pre class="small text-info m-0">{{ formatJSON(resultadoVeriphone) }}</pre>
              </div>
            </div>
          </div>
        </div>

        <!-- IPQS -->
        <div class="col-md-5">
          <div class="card h-100 bg-dark border-secondary shadow-sm">
            <div class="card-header border-secondary d-flex justify-content-between">
              <span class="fw-bold">IPQualityScore</span>
              <span :class="ipqsKey ? 'text-success' : 'text-danger'">{{ ipqsKey ? 'Ready' : 'Missing Key' }}</span>
            </div>
            <div class="card-body d-flex flex-column">
              <button class="btn btn-sm btn-primary w-100 mb-3" :disabled="!ipqsKey || !targetPhone || loadingIpqs" @click="ejecutarIpqs">
                <span v-if="loadingIpqs" class="spinner-border spinner-border-sm me-2"></span>
                Execute IPQS
              </button>
              <div v-if="resultadoIpqs" class="flex-grow-1 bg-secondary rounded p-2 overflow-auto" style="max-height: 250px;">
                <pre class="small text-info m-0">{{ formatJSON(resultadoIpqs) }}</pre>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Quantum Report -->
      <div class="text-center">
        <button class="btn btn-lg btn-success px-5" :disabled="!canAnalyze" @click="generarReporteQuantum">
          <span v-if="aiLoading" class="spinner-border spinner-border-sm me-2"></span>
          Reporte Completo IA
        </button>
        <div v-if="aiReport" class="mt-4 p-4 text-start bg-dark border border-info rounded shadow">
          <pre class="text-light m-0" style="white-space: pre-wrap; font-family: monospace;">{{ aiReport }}</pre>
        </div>
      </div>
    </main>

    <!-- BYOK Modal -->
    <div class="modal fade" id="configModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content bg-dark text-light border-secondary">
          <div class="modal-header border-secondary">
            <h5 class="modal-title text-info">API Keys Configuration</h5>
            <button type="button" class="btn-close btn-close-white" data-bs-dismiss="modal"></button>
          </div>
          <div class="modal-body">
            <div class="mb-3">
              <label class="form-label">Groq API Key</label>
              <input type="password" class="form-control bg-secondary text-white border-dark" v-model="groqKey">
              <a href="https://console.groq.com/keys" target="_blank" class="text-secondary small text-decoration-none mt-1 d-block">🔗 Consigue tu API aquí</a>
            </div>
            <div class="mb-3">
              <label class="form-label">Veriphone API Key</label>
              <input type="password" class="form-control bg-secondary text-white border-dark" v-model="veriphoneKey">
              <a href="https://veriphone.io/" target="_blank" class="text-secondary small text-decoration-none mt-1 d-block">🔗 Consigue tu API aquí</a>
            </div>
            <div class="mb-3">
              <label class="form-label">IPQualityScore API Key</label>
              <input type="password" class="form-control bg-secondary text-white border-dark" v-model="ipqsKey">
              <a href="https://www.ipqualityscore.com/create-account" target="_blank" class="text-secondary small text-decoration-none mt-1 d-block">🔗 Consigue tu API aquí</a>
            </div>
          </div>
          <div class="modal-footer border-secondary">
            <button type="button" class="btn btn-success" data-bs-dismiss="modal" @click="saveKeys">Save & Reload</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const targetPhone = ref('')
const aiLoading = ref(false)
const aiReport = ref('')

const groqKey = ref('')
const veriphoneKey = ref('')
const ipqsKey = ref('')

const resultadoVeriphone = ref(null)
const resultadoIpqs = ref(null)

const loadingVeriphone = ref(false)
const loadingIpqs = ref(false)

const apiBase = 'http://localhost:5000/api'

onMounted(() => {
  groqKey.value = localStorage.getItem('groqKey') || ''
  veriphoneKey.value = localStorage.getItem('veriphoneKey') || ''
  ipqsKey.value = localStorage.getItem('ipqsKey') || ''
})

const saveKeys = () => {
  localStorage.setItem('groqKey', groqKey.value)
  localStorage.setItem('veriphoneKey', veriphoneKey.value)
  localStorage.setItem('ipqsKey', ipqsKey.value)
}

const formatJSON = (obj) => JSON.stringify(obj, null, 2)

const canAnalyze = computed(() => {
  return groqKey.value && targetPhone.value && (resultadoVeriphone.value || resultadoIpqs.value)
})

const ejecutarVeriphone = async () => {
  loadingVeriphone.value = true
  resultadoVeriphone.value = null
  try {
    const res = await axios.get(`${apiBase}/scan/veriphone?target=${encodeURIComponent(targetPhone.value)}`, {
      headers: { 'X-Veriphone-Key': veriphoneKey.value }
    })
    resultadoVeriphone.value = res.data
  } catch (err) {
    resultadoVeriphone.value = { error: err.message, details: err.response?.data }
  } finally {
    loadingVeriphone.value = false
  }
}

const ejecutarIpqs = async () => {
  loadingIpqs.value = true
  resultadoIpqs.value = null
  try {
    const res = await axios.get(`${apiBase}/scan/ipqs?target=${encodeURIComponent(targetPhone.value)}`, {
      headers: { 'X-Ipqs-Key': ipqsKey.value }
    })
    resultadoIpqs.value = res.data
  } catch (err) {
    resultadoIpqs.value = { error: err.message, details: err.response?.data }
  } finally {
    loadingIpqs.value = false
  }
}

const generarReporteQuantum = async () => {
  aiLoading.value = true
  aiReport.value = ''
  
  const payload = {
    messages: [{
      role: 'user',
      content: `REGLA ABSOLUTA: PROHIBIDO usar formato Markdown. NO uses asteriscos (**), ni numerales (#). Escribe el reporte como un humano, en texto plano, usando párrafos normales y guiones simples (-) para las listas. Sé directo y profesional. Datos a analizar: ${JSON.stringify({veriphone: resultadoVeriphone.value, ipqs: resultadoIpqs.value})}`
    }],
    model: 'llama-3.3-70b-versatile'
  }

  try {
    const res = await axios.post('https://api.groq.com/openai/v1/chat/completions', payload, {
      headers: {
        'Authorization': `Bearer ${groqKey.value}`,
        'Content-Type': 'application/json'
      }
    })
    aiReport.value = res.data.choices[0].message.content
  } catch (err) {
    aiReport.value = `Analysis Failed: ${err.message}`
  } finally {
    aiLoading.value = false
  }
}
</script>
