<!-- ManagerDashboard.vue -->
<template>
    <div class="manager-dashboard container mx-auto px-6 py-8">
      <div class="flex justify-between items-center mb-6">
        <h2>Manager Dashboard</h2>
        <div class="relative">
          <button 
            @click="isDropdownOpen = !isDropdownOpen"
            class="px-4 py-2 text-sm bg-blue-500 text-white rounded-md hover:bg-blue-600"
          >
            Connect with
          </button>
          <div 
            v-if="isDropdownOpen" 
            class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg z-10"
          >
            <a 
              v-for="service in services" 
              :key="service"
              href="#"
              @click.prevent="connectWith(service)"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
            >
              {{ service }}
            </a>
          </div>
        </div>
      </div>
      
      <!-- Loading State -->
      <div v-if="onboardingStore.isLoading" class="loading">
        Generating onboarding checklist...
      </div>
  
      <!-- Error State -->
      <div v-if="onboardingStore.error" class="error">
        {{ onboardingStore.error }}
      </div>
  
      <!-- Upload Form -->
      <form @submit.prevent="handleSubmit" v-if="!onboardingStore.isLoading" class="space-y-6 max-w-lg mx-auto bg-white p-6 rounded-lg shadow-sm">
        <div class="space-y-4">
          <input 
            v-model="engineerEmail" 
            type="email" 
            placeholder="New Hire Email" 
            required
            class="w-full"
          >
          <input 
            type="file" 
            @change="handleFileUpload" 
            accept=".txt"
            required
            class="w-full"
          >
        </div>
        <button 
          type="submit"
          class="px-4 py-2 text-sm bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          Generate Onboarding
        </button>
      </form>
  
      <!-- Success State -->
      <div v-if="onboardingStore.hasOnboardingData" class="success">
        Onboarding checklist generated for {{ onboardingStore.engineerEmail }}
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useOnboardingStore } from '../stores/onboarding'
  
  const onboardingStore = useOnboardingStore()
  const isDropdownOpen = ref(false)
  const services = ['Notion', 'Jira', 'GitHub', 'GitLab']

  const connectWith = (service) => {
    isDropdownOpen.value = false
    // TODO: Implement connection logic
    console.log(`Connecting with ${service}`)
  }
  const engineerEmail = ref('')
  const documentContent = ref('')
  
  const handleFileUpload = (event) => {
    const file = event.target.files[0]
    const reader = new FileReader()
    reader.onload = (e) => {
      documentContent.value = e.target.result
    }
    reader.readAsText(file)
  }
  
  const handleSubmit = async () => {
    await onboardingStore.generateOnboarding(
      engineerEmail.value, 
      documentContent.value
    )
  }
  </script>
  
  