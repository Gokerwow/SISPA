<script setup lang="ts">
import { ref } from 'vue'
import { AddServices } from '@/services/services'
import type { ServicesCreate } from '@/types/services'
import { useToast } from 'vue-toastification'
import { useRouter } from 'vue-router'
import ServiceForm from '@/components/service/serviceForm.vue'

const isSubmitting = ref(false)
const toast = useToast()
const router = useRouter()

// Submit form
const handleSubmit = async (formData: ServicesCreate) => {
  isSubmitting.value = true
  const { name, price_male, price_female, price_tourist, is_active } = formData
  const cleanPayload = { name, price_male, price_female, price_tourist, is_active }

  try {
    const result = await AddServices(cleanPayload)

    console.log('Form submitted:', formData, result)
    toast.success('Successfully Created The Service')
    router.push('/services')
  } catch (error) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const message = (error as any).response?.data?.error || "Failed to add service";
    console.error('Error deleting service:', message)
    toast.error(message);
  } finally {
    isSubmitting.value = false
  }
}

</script>

<template>
  <ServiceForm @submit="handleSubmit"/>
</template>
