<script setup lang="ts">
import { GetCustomerDetail, updateCustomer } from '@/services/customers';
import type { Customer } from '@/types/customer';
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import CustomerForm from '@/components/customer/customerForm.vue';
import { useToast } from 'vue-toastification';

const customer = ref<Customer | null>(null)

const isLoading = ref(true)
const route = useRoute()
const router = useRouter()
const toast = useToast()

const fetchCustomer = async (id: number) => {
  isLoading.value = true
  try {
    const response = await GetCustomerDetail(id)
    customer.value = response
    console.log(response)
  } catch (error) {
    console.log(error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  const id = Number(route.params.id)
  if (id) {
    fetchCustomer(id)
  }
})

const handleUpdate = async (editedData: Customer) => {
  try {
    const id = Number(route.params.id)
    const result = await updateCustomer(id, editedData)
    console.log(result)
    toast.success('Successfully updated customer data')
    router.push('/customers')
  } catch (error) {
    console.log(error)
    toast.error(`error updating customer data: ${error}`)
  }
}

</script>

<template>
  <div v-if="customer">
    <CustomerForm :initial-data="customer" @submit="handleUpdate" />
  </div>
  <div v-else>
    <p>Loading customer profile...</p>
  </div>
</template>

<style scoped></style>
