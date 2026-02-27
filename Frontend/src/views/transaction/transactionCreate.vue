<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import { GetAllCustomers } from '@/services/customers'
import { GetAllTherapist } from '@/services/therapist'
import { GetAllServices } from '@/services/services'
import { AddTransaction } from '@/services/transactions'
import type { TransactionRequest } from '@/types/transaction'
import TransactionForm from '@/components/transaction/transactionForm.vue'
import type { FullCustomer } from '@/types/customer'
import type { Therapist } from '@/types/therapist'
import type { Services } from '@/types/services'
import { validateForm } from '@/utils/validateForm'

const router = useRouter()
const toast = useToast()

const isSubmitting = ref(false)

const customers = ref<FullCustomer[]>([])
const therapists = ref<Therapist[]>([])
const services = ref<Services[]>([])

const loadingCustomers = ref(false)
const loadingTherapists = ref(false)
const loadingServices = ref(false)

// Fetch data
// Fetch data
const fetchCustomers = async () => {
  loadingCustomers.value = true
  try {
    const response = await GetAllCustomers()
    customers.value = response.data
  } catch (error) {
    console.error('Error fetching customers:', error)
  } finally {
    loadingCustomers.value = false
  }
}

const fetchTherapists = async () => {
  loadingTherapists.value = true
  try {
    const response = await GetAllTherapist()
    therapists.value = response.data
  } catch (error) {
    console.error('Error fetching therapists:', error)
  } finally {
    loadingTherapists.value = false
  }
}

const fetchServices = async () => {
  loadingServices.value = true
  try {
    const response = await GetAllServices()
    services.value = response.data
  } catch (error) {
    console.error('Error fetching services:', error)
  } finally {
    loadingServices.value = false
  }
}

// Load data on mount
fetchCustomers()
fetchTherapists()
fetchServices()

// Submit form
const handleSubmit = async (formData: TransactionRequest) => {
  if (!validateForm(formData)) return

  isSubmitting.value = true

  try {
    const response = await AddTransaction(formData)

    const result = response
    console.log('Transaction created:', result)

    toast.success('Transaction created successfully!')
    router.push('/transactions')
  } catch (error) {
    const message = (error as Error).message || "Failed to create transaction"
    console.error('Error creating transaction:', message)
    toast.error(message)
  } finally {
    isSubmitting.value = false
  }
}

</script>

<template>
  <TransactionForm @submit="handleSubmit" :initial-data-customer="customers" :initial-data-service="services" :initial-data-therapist="therapists" :is-submitting="isSubmitting"/>
</template>

