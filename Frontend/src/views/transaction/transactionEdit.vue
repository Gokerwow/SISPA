<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import { GetAllCustomers } from '@/services/customers'
import { GetAllTherapist } from '@/services/therapist'
import { GetAllServices } from '@/services/services'
import { GetInvoiceDetails, updateTransaction } from '@/services/transactions'
import type { InvoiceData, TransactionRequest } from '@/types/transaction'
import TransactionForm from '@/components/transaction/transactionForm.vue'
import type { FullCustomer } from '@/types/customer'
import type { Therapist } from '@/types/therapist'
import type { Services } from '@/types/services'
import { validateForm } from '@/utils/validateForm'

const router = useRouter()
const route = useRoute()
const toast = useToast()

const isSubmitting = ref(false)

const invoice = ref<InvoiceData>()
const customers = ref<FullCustomer[]>([])
const therapists = ref<Therapist[]>([])
const services = ref<Services[]>([])
const id = Number(route.params.id)

const loadingInvoice = ref(false)
const loadingCustomers = ref(false)
const loadingTherapists = ref(false)
const loadingServices = ref(false)

// Fetch data
const fetchInvoice = async (id: number) => {
  loadingInvoice.value = true
  try {
    const response = await GetInvoiceDetails(id)
    invoice.value = response.data
    console.log(invoice.value)
  } catch (error) {
    console.error('Error fetching customers:', error)
  } finally {
    loadingInvoice.value = false
  }
}

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

onMounted(() => {
  if (id) {
    fetchInvoice(id)
  }
})

// Load data on mount
fetchCustomers()
fetchTherapists()
fetchServices()

// Submit form
const handleSubmit = async (formData: TransactionRequest) => {
  if (!validateForm(formData)) return

  isSubmitting.value = true

  try {
    const response = await updateTransaction(id, formData)

    const result = response
    console.log('Transaction Updated:', result)

    toast.success('Transaction Updated successfully!')
    router.push('/transactions')
  } catch (error) {
    const message = (error as Error).message || "Failed to Updated transaction"
    console.error('Error Updating transaction:', message)
    toast.error(message)
  } finally {
    isSubmitting.value = false
  }
}

</script>

<template>
  <TransactionForm @submit="handleSubmit" :initial-data-customer="customers" :initial-data-service="services" :initial-data-therapist="therapists" :initial-data-invoice="invoice" :is-submitting="isSubmitting"/>
</template>

