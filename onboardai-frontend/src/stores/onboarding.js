import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useOnboardingStore = defineStore('onboarding', () => {
  // State
  const checklist = ref(null)
  const codingChallenge = ref(null)
  const isLoading = ref(false)
  const error = ref(null)
  const engineerEmail = ref('')
  const uploadedDocument = ref(null)

  // Computed
  const hasOnboardingData = computed(() => {
    return checklist.value !== null && codingChallenge.value !== null
  })

  const progress = computed(() => {
    if (!checklist.value) return 0
    const completedItems = checklist.value.filter(item => item.completed).length
    return (completedItems / checklist.value.length) * 100
  })

  // Actions
  const generateOnboarding = async (email, documentation) => {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await fetch('http://localhost:8080/api/generate-onboarding', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          engineerEmail: email,
          documentation
        })
      })

      if (!response.ok) {
        throw new Error(await response.text() || 'Failed to generate onboarding')
      }

      const data = await response.json()
      // Transform checklist items to include completion status
      checklist.value = data.checklist.map(item => ({
        text: item,
        completed: false
      }))
      codingChallenge.value = data.codingChallenge
      engineerEmail.value = email
      uploadedDocument.value = documentation
      
      return true
    } catch (e) {
      error.value = e.message
      return false
    } finally {
      isLoading.value = false
    }
  }

  const fetchEngineerOnboarding = async (email) => {
    if (!email) return false
    
    isLoading.value = true
    error.value = null

    try {
      const response = await fetch(`http://localhost:8080/api/onboarding/${email}`)
      
      if (!response.ok) {
        throw new Error('No onboarding data found')
      }

      const data = await response.json()
      checklist.value = data.checklist.map(item => ({
        text: item,
        completed: false
      }))
      codingChallenge.value = data.codingChallenge
      engineerEmail.value = email
      return true
    } catch (e) {
      error.value = e.message
      return false
    } finally {
      isLoading.value = false
    }
  }

  const toggleItemCompletion = (index) => {
    if (checklist.value && checklist.value[index]) {
      checklist.value[index].completed = !checklist.value[index].completed
    }
  }

  const clearOnboardingData = () => {
    checklist.value = null
    codingChallenge.value = null
    engineerEmail.value = ''
    uploadedDocument.value = null
    error.value = null
  }

  return {
    // State
    checklist,
    codingChallenge,
    isLoading,
    error,
    engineerEmail,
    uploadedDocument,

    // Computed
    hasOnboardingData,
    progress,

    // Actions
    generateOnboarding,
    fetchEngineerOnboarding,
    toggleItemCompletion,
    clearOnboardingData
  }
})