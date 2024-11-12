<template>
    <div class="max-w-2xl mx-auto p-6 bg-white rounded-lg shadow-md">
      <div v-if="isLoading" class="flex justify-center items-center py-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
      </div>
  
      <div v-else-if="error" class="text-center py-8">
        <p class="text-red-500">{{ error }}</p>
      </div>
  
      <template v-else>
        <!-- Checklist Section -->
        <div class="mb-8">
          <h3 class="text-lg font-semibold mb-4">Onboarding Checklist</h3>
          <div v-if="checklist && checklist.length" class="space-y-2">
            <div
              v-for="(item, index) in checklist"
              :key="index"
              class="flex items-start space-x-3 p-2 hover:bg-gray-50 rounded"
            >
              <input
                type="checkbox"
                :id="'item-' + index"
                v-model="item.completed"
                @change="onToggleItem(index)"
                class="mt-1"
              >
              <label 
                :for="'item-' + index" 
                class="text-gray-700 cursor-pointer"
                :class="{ 'line-through text-gray-400': item.completed }"
              >
                {{ item.text }}
              </label>
            </div>
          </div>
          <p v-else class="text-gray-500 text-center py-4">
            No checklist items available
          </p>
        </div>
  
        <!-- Coding Challenge Section -->
        <div v-if="codingChallenge" class="mt-6">
          <h3 class="text-lg font-semibold mb-4">Coding Challenge</h3>
          <div class="bg-gray-50 p-4 rounded-md">
            <pre class="whitespace-pre-wrap text-sm">{{ codingChallenge }}</pre>
          </div>
        </div>
  
        <!-- Progress Section -->
        <div v-if="checklist && checklist.length" class="mt-6">
          <div class="flex justify-between items-center mb-2">
            <span class="text-sm font-medium text-gray-700">Progress</span>
            <span class="text-sm text-gray-500">
              {{ completedCount }} / {{ checklist.length }} completed
            </span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2.5">
            <div
              class="bg-blue-500 h-2.5 rounded-full transition-all duration-300"
              :style="{ width: `${progress}%` }"
            ></div>
          </div>
        </div>
      </template>
    </div>
  </template>
  
  <script setup>
  import { computed } from 'vue'
  import { useOnboardingStore } from '../stores/onboarding'
  
  const props = defineProps({
    checklist: {
      type: Array,
      default: () => []
    },
    codingChallenge: {
      type: String,
      default: ''
    },
    isLoading: {
      type: Boolean,
      default: false
    },
    error: {
      type: String,
      default: ''
    }
  })
  
  const onboardingStore = useOnboardingStore()
  
  const completedCount = computed(() => {
    return props.checklist.filter(item => item.completed).length
  })
  
  const progress = computed(() => {
    if (!props.checklist.length) return 0
    return (completedCount.value / props.checklist.length) * 100
  })
  
  const emit = defineEmits(['toggle-item'])
  
  const onToggleItem = (index) => {
    emit('toggle-item', index)
  }
  </script>