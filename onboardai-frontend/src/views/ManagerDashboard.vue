<!-- ManagerDashboard.vue -->
<template>
    <div class="manager-dashboard">
      <h2>Manager Dashboard</h2>
      
      <!-- Loading State -->
      <div v-if="onboardingStore.isLoading" class="loading">
        Generating onboarding checklist...
      </div>
  
      <!-- Error State -->
      <div v-if="onboardingStore.error" class="error">
        {{ onboardingStore.error }}
      </div>
  
      <!-- Upload Form -->
      <form @submit.prevent="handleSubmit" v-if="!onboardingStore.isLoading">
        <input 
          v-model="engineerEmail" 
          type="email" 
          placeholder="New Hire Email" 
          required
        >
        <input 
          type="file" 
          @change="handleFileUpload" 
          accept=".txt"
          required
        >
        <button type="submit">Generate Onboarding</button>
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
  
  