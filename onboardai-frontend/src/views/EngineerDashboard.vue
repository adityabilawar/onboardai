<template>
    <div class="container mx-auto px-4 py-8">
      <div class="max-w-4xl mx-auto">
        <div class="mb-8">
          <h1 class="text-2xl font-bold">Engineer Dashboard</h1>
        </div>
  
        <div v-if="userStore.error" class="mb-6 p-4 bg-red-100 text-red-700 rounded">
          {{ userStore.error }}
        </div>
  
        <div v-if="onboardingStore.isLoading" class="text-center py-12">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
          <p class="mt-4 text-gray-600">Loading your onboarding checklist...</p>
        </div>
  
        <template v-else>
          <Checklist
            v-if="onboardingStore.hasOnboardingData"
            :checklist="onboardingStore.checklist || []"
            :coding-challenge="onboardingStore.codingChallenge"
            :is-loading="onboardingStore.isLoading"
            :error="onboardingStore.error"
            @toggle-item="onboardingStore.toggleItemCompletion"
          />
  
          <div 
            v-else 
            class="text-center py-12 bg-gray-50 rounded-lg"
          >
            <p class="text-gray-600">
              Waiting for your manager to upload documentation and generate your onboarding checklist.
            </p>
          </div>
        </template>
      </div>
    </div>
  </template>
  
  <script setup>
  import { onMounted } from 'vue'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '../stores/user'
  import { useOnboardingStore } from '../stores/onboarding'
  import Checklist from '../components/Checklist.vue'
  
  const router = useRouter()
  const userStore = useUserStore()
  const onboardingStore = useOnboardingStore()
  
  onMounted(async () => {
    if (!userStore.isAuthenticated) {
      router.push('/')
      return
    }
  
    if (userStore.userEmail) {
      await onboardingStore.fetchEngineerOnboarding(userStore.userEmail)
    }
  })
  
  const handleLogout = () => {
    userStore.logout()
    onboardingStore.clearOnboardingData()
    router.push('/')
  }
  </script>