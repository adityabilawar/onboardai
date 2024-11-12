<template>
  <div class="min-h-screen bg-gray-50 font-sans">
    <!-- Loading bar -->
    <div v-if="userStore.isLoading || onboardingStore.isLoading" class="fixed top-0 left-0 right-0">
      <div class="h-1 bg-blue-500 animate-[loading_2s_ease-in-out_infinite]"></div>
    </div>
    <!-- Navigation when user is authenticated -->
    <nav v-if="userStore.isAuthenticated" class="bg-white shadow">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <h1 class="text-xl font-bold text-gray-900">Onboard AI</h1>
            </div>
          </div>
          <div class="flex items-center">
            <span class="text-gray-600 mr-4">{{ userStore.userEmail }}</span>
            <button
              @click="logout"
              class="px-4 py-2 text-sm text-red-600 border border-red-600 rounded hover:bg-red-50"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main content -->
    <main>
      <RouterView />
    </main>
  </div>
</template>

<script setup>
import { RouterView, useRouter } from 'vue-router'
import { useUserStore } from './stores/user'
import { useOnboardingStore } from './stores/onboarding'

const router = useRouter()
const userStore = useUserStore()
const onboardingStore = useOnboardingStore()

const logout = () => {
  userStore.logout()
  onboardingStore.clearOnboardingData()
  router.push('/')
}
</script>