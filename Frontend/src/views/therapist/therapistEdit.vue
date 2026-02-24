<script setup lang="ts">
import TherapistForm from '@/components/therapist/therapistForm.vue'
import { GetTherapistDetail, updateTherapist } from '@/services/therapist'
import type { Therapist } from '@/types/therapist'
import { formatForInput } from '@/utils/date'
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'

const therapist = ref<Therapist | undefined>(undefined)

const isLoading = ref(true)
const route = useRoute()
const router = useRouter()
const toast = useToast()

const fetchTherapist = async (id: number) => {
  try {
    const response = await GetTherapistDetail(id)
    therapist.value = response
    if (therapist.value) {
      therapist.value.joined_date = formatForInput(therapist.value.joined_date) ?? null
    }
    console.log(therapist.value)
  } catch (error) {
    console.log(error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  const id = Number(route.params.id)
  if (id) {
    fetchTherapist(id)
  }
})

const handleUpdate = async (editedData: Therapist) => {
  try {
    const id = Number(route.params.id)
    const result = await updateTherapist(id, editedData)
    console.log(result)
    toast.success('Successfully updated therapist data')
    router.push('/therapists')
  } catch (error) {
    console.log(error)
    toast.error(`error updating therapist data: ${error}`)
  }
}
</script>

<template>
  <div v-if="therapist">
    <TherapistForm :initial-data="therapist" @submit="handleUpdate"/>
  </div>
  <div v-else>
    <p>Loading therapist details...</p>
  </div>
</template>

<style scoped></style>
