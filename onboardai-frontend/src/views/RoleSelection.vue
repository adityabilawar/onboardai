<template>
    <div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div class="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Select Your Role
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          Are you a manager or a new hire engineer?
        </p>
      </div>
  
      <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
          <div v-if="userStore.error" class="mb-4 p-3 bg-red-100 text-red-700 rounded">
            {{ userStore.error }}
          </div>
  
          <div class="space-y-4">
            <button
              @click="selectRole('manager')"
              class="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              :disabled="userStore.isLoading"
            >
              I'm a Manager
            </button>
  
            <button
              @click="selectRole('engineer')"
              class="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
              :disabled="userStore.isLoading"
            >
              I'm a New Hire Engineer
            </button>
          </div>
  
          <div class="mt-6">
            <button
              @click="handleLogout"
              class="w-full flex justify-center py-2 px-4 text-sm text-gray-600 hover:text-gray-900"
            >
              Back to Login
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { onMounted } from 'vue'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '../stores/user'
  
  const router = useRouter()
  const userStore = useUserStore()
  
  onMounted(() => {
    if (!userStore.isAuthenticated) {
      router.push('/')
    }
  })
  
  const selectRole = async (role) => {
    userStore.setRole(role)
    router.push(role === 'manager' ? '/manager-dashboard' : '/engineer-dashboard')
  }
  
  const handleLogout = () => {
    userStore.logout()
    router.push('/')
  }
  </script>