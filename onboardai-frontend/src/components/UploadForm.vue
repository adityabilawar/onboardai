<!-- src/components/UploadForm.vue -->
<template>
    <form @submit.prevent="handleSubmit" class="max-w-md mx-auto p-6 bg-white rounded-lg shadow-md">
      <div class="mb-4">
        <label for="engineerEmail" class="block text-sm font-medium text-gray-700 mb-1">
          Engineer's Email
        </label>
        <input
          id="engineerEmail"
          v-model="engineerEmail"
          type="email"
          required
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="{ 'border-red-500': errors.email }"
          placeholder="Enter engineer's email"
        >
        <p v-if="errors.email" class="mt-1 text-sm text-red-500">{{ errors.email }}</p>
      </div>
  
      <div class="mb-6">
        <label class="block text-sm font-medium text-gray-700 mb-1">
          Documentation File
        </label>
        <div class="mt-1 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-md">
          <div class="space-y-1 text-center">
            <div class="flex flex-col items-center">
              <input
                type="file"
                @change="handleFileChange"
                accept=".txt,.md,.doc,.docx"
                class="sr-only"
                ref="fileInput"
              >
              <button
                type="button"
                @click="$refs.fileInput.click()"
                class="px-4 py-2 text-sm text-blue-500 border border-blue-500 rounded-md hover:bg-blue-50"
              >
                Select File
              </button>
            </div>
            <p class="text-xs text-gray-500">
              {{ fileName || 'Upload documentation file' }}
            </p>
          </div>
        </div>
        <p v-if="errors.file" class="mt-1 text-sm text-red-500">{{ errors.file }}</p>
      </div>
  
      <div class="flex justify-end">
        <button
          type="submit"
          class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          :disabled="isLoading"
        >
          {{ isLoading ? 'Generating...' : 'Generate Onboarding' }}
        </button>
      </div>
    </form>
  </template>
  
  <script setup>
  import { ref, reactive } from 'vue'
  import { useOnboardingStore } from '../stores/onboarding'
  
  const onboardingStore = useOnboardingStore()
  
  const engineerEmail = ref('')
  const fileInput = ref(null)
  const fileName = ref('')
  const isLoading = ref(false)
  const errors = reactive({
    email: '',
    file: ''
  })
  
  const handleFileChange = (event) => {
    const file = event.target.files[0]
    if (file) {
      fileName.value = file.name
      const reader = new FileReader()
      reader.onload = (e) => {
        documentContent.value = e.target.result
      }
      reader.readAsText(file)
    }
  }
  
  const validateForm = () => {
    let isValid = true
    errors.email = ''
    errors.file = ''
  
    if (!engineerEmail.value) {
      errors.email = 'Engineer email is required'
      isValid = false
    }
    if (!documentContent.value) {
      errors.file = 'Documentation file is required'
      isValid = false
    }
  
    return isValid
  }
  
  const handleSubmit = async () => {
    if (!validateForm()) return
  
    isLoading.value = true
    try {
      await onboardingStore.generateOnboarding(engineerEmail.value, documentContent.value)
    } catch (error) {
      errors.file = 'Failed to generate onboarding'
    } finally {
      isLoading.value = false
    }
  }
  
  const documentContent = ref('')
  </script>