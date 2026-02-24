<script setup lang="ts">
import { ref } from 'vue'
import { useToast } from 'vue-toastification'
import { useRouter } from 'vue-router'
import TherapistForm from '@/components/therapist/therapistForm.vue'
import type { Therapist } from '@/types/therapist'
import { AddTherapist } from '@/services/therapist'

const isSubmitting = ref(false)
const toast = useToast()
const router = useRouter()

// Submit form
const handleSubmit = async (formData: Therapist) => {
  isSubmitting.value = true
  const { name, phone, address, commission_rate, joined_date, exit_date, status } = formData
  const cleanPayload = { name, phone, address, commission_rate, joined_date, exit_date, status }

  try {
    const result = await AddTherapist(cleanPayload)

    console.log('Therapist created:', result)

    toast.success('Successfully Created The Therapist')
    router.push('/therapists')
  } catch (error) {
    const message = (error as any).response?.data?.error || (error as Error).message || "Failed to add therapist"
    console.error('Error adding therapist:', message)
    toast.error(message)
  } finally {
    isSubmitting.value = false
  }
}

</script>

<template>
  <TherapistForm @submit="handleSubmit" />
</template>
