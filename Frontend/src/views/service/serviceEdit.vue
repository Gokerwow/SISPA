<script setup lang="ts">
import { GetServiceDetail, updateService } from '@/services/services'
import type { Services } from '@/types/services'
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import ServiceForm from '@/components/service/serviceForm.vue'


const service = ref<Services | undefined>(undefined)

const isLoading = ref(true)
const route = useRoute()
const router = useRouter()
const toast = useToast()
const id = Number(route.params.id)

const fetchService = async (id: number) => {
  try {
    const response = await GetServiceDetail(id)
    service.value = response.data
    console.log(service.value)
  } catch (error) {
    console.log(error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  if (id) {
    fetchService(id)
  }
})

const handleUpdate = async (editedData: Services) => {
  try {
    const id = Number(route.params.id)
    const result = await updateService(id, editedData)
    console.log(result)
    toast.success('Successfully updated customer data')
    router.push('/services')
  } catch (error) {
    console.log(error)
    toast.error(`error updating customer data: ${error}`)
  }
}

</script>

<template>
  <div v-if="service">
    <ServiceForm :initial-data="service" @submit="handleUpdate" />
  </div>
  <div v-else>
    <p>Loading service details...</p>
  </div>
</template>

<style scoped></style>
