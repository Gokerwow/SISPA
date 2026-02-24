<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { AddCustomers } from '@/services/customers'
import { useToast } from 'vue-toastification'
import CustomerForm from '@/components/customer/customerForm.vue'
import type { Customer } from '@/types/customer'

const isSubmitting = ref(false)
const router = useRouter()
const toast = useToast();

// Submit form
const handleSubmit = async (formData: Customer) => {
  isSubmitting.value = true
  const { name, phone, birth_date, gender, address, note } = formData
  const cleanPayload = { name, phone, birth_date, gender, address, note }

  try {
    const result = await AddCustomers(cleanPayload)
    console.log('Form submitted:', formData)
    console.log(result)
    toast.success("Customer added successfully!");
    router.push('/customers')
  } catch (error) {
    const message = error.response?.data?.error || "Failed to add customer";
    console.error('Error submitting form:', message)
    toast.error(message);
  } finally {
    isSubmitting.value = false
  }
}

</script>

<template>
  <CustomerForm @submit="handleSubmit"/>
</template>
