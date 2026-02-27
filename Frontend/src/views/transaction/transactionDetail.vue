<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  FileText,
  Calendar,
  DollarSign,
  ArrowLeft,
  Edit2,
  Users,
  UserCircle,
  Heart,
  ShoppingCart,
  Printer,
} from 'lucide-vue-next'
import { GetInvoiceDetails } from '@/services/transactions'
import type { InvoiceData } from '@/types/transaction'
import { GetCustomerDetail } from '@/services/customers'
import { GetServiceDetail } from '@/services/services'
import { GetTherapistDetail } from '@/services/therapist'

const route = useRoute()
const router = useRouter()
const id = Number(route.params.id)

const invoice = ref<InvoiceData | null>(null)
const loading = ref(true)
const error = ref("")

// Mock data for related entities (replace with actual API calls)
const customers = ref<Record<number, { name: string, phone: string }>>({})
const therapists = ref<Record<number, { name: string }>>({})
const services = ref<Record<number, { name: string }>>({})
// const customers = ref<Customer[]>([])
// const therapists = ref<Therapist[]>([])
// const services = ref<Services[]>([])

// Fetch invoice data
const fetchInvoice = async (id: number) => {
  try {
    const response = await GetInvoiceDetails(id)
    invoice.value = response.data
    console.log(invoice.value)

    const fetchTasks = invoice.value?.items.map(item => {
      return fetchRelatedData(item.customer_id, item.service_id, item.therapist_id);
    }) || [];

    await Promise.all(fetchTasks);
    console.log(customers.value, services.value, therapists.value)

    // Fetch related data
  } catch (err) {
    error.value = (err as Error).message
    console.error('Error fetching invoice:', err)
  } finally {
    loading.value = false
  }
}

// Fetch customers, therapists, and services
const fetchRelatedData = async (cusID: number, servID: number, therID: number) => {
  try {
    // Fetch customers
    const customerRes = await GetCustomerDetail(cusID)
    customers.value[cusID] = customerRes.data

    // Fetch therapists
    const therapistsRes = await GetTherapistDetail(therID)
    therapists.value[therID] = therapistsRes.data

    // Fetch services
    const servicesRes = await GetServiceDetail(servID)
    services.value[servID] = servicesRes.data

  } catch (error) {
    console.error('Error fetching related data:', error)
  }
}

onMounted(() => {
  if (id) {
    fetchInvoice(id)
  }
})

// Format currency
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// Format date
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('id-ID', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

// Get status badge color
const getStatusColor = (status: string) => {
  const colors: Record<string, string> = {
    'paid': 'bg-emerald-100 text-emerald-700 border-emerald-300',
    'pending': 'bg-amber-100 text-amber-700 border-amber-300',
    'cancelled': 'bg-red-100 text-red-700 border-red-300'
  }
  return colors[status?.toLowerCase()] || 'bg-slate-100 text-slate-700 border-slate-300'
}

// Get item gradient
const getItemGradient = (id: number) => {
  const gradients = [
    'from-purple-400 to-indigo-500',
    'from-rose-400 to-pink-500',
    'from-emerald-400 to-teal-500',
    'from-orange-400 to-red-500',
    'from-blue-400 to-cyan-500',
    'from-fuchsia-400 to-purple-500'
  ]
  return gradients[id % gradients.length]
}

// Calculate therapist commission
// const calculateCommission = (price: number, therapistId: number) => {
//   const therapist = therapists.value[therapistId]
//   if (!therapist) return 0
//   return (price * therapist.commission_rate) / 100
// }

// Print invoice
const printInvoice = () => {
  window.print()
}
</script>

<template>
  <div class="space-y-6">

    <!-- Back Button -->
    <button
      @click="router.back()"
      class="flex items-center gap-2 text-slate-600 hover:text-slate-800 transition-colors cursor-pointer">
      <ArrowLeft class="w-4 h-4" />
      <span class="text-sm font-medium">Back to Transactions</span>
    </button>

    <!-- Loading State -->
    <div v-if="loading" class="bg-white/80 backdrop-blur-xl rounded-2xl border border-white/60 shadow-lg p-20">
      <div class="text-center">
        <div class="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gradient-to-r from-rose-400 via-purple-400 to-indigo-400 mb-4 animate-pulse">
          <FileText class="w-8 h-8 text-white" />
        </div>
        <p class="text-slate-500 font-medium">Loading transaction details...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-50 border-l-4 border-red-500 rounded-xl p-5 shadow-lg">
      <div class="flex items-start gap-3">
        <div class="w-10 h-10 rounded-xl bg-red-100 flex items-center justify-center flex-shrink-0">
          <span class="text-red-600 text-xl">⚠️</span>
        </div>
        <div>
          <p class="font-bold text-red-800 mb-1">Error Loading Transaction</p>
          <p class="text-red-700 text-sm">{{ error }}</p>
        </div>
      </div>
    </div>

    <!-- Invoice Details -->
    <div v-else-if="invoice" class="space-y-6">

      <!-- Header Section -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-6">
          <div class="flex items-start gap-4">
            <div class="w-16 h-16 rounded-2xl bg-gradient-to-br from-purple-400 to-indigo-500 flex items-center justify-center shadow-lg">
              <FileText class="w-8 h-8 text-white" />
            </div>
            <div>
              <h1 class="text-3xl font-bold text-slate-800 mb-1">{{ invoice.nota_number }}</h1>
              <p class="text-slate-500 text-sm">Invoice #{{ invoice.id }}</p>
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="flex items-center gap-2">
            <button
              @click="printInvoice"
              class="px-4 py-2 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-xl text-sm font-medium transition-all flex items-center gap-2">
              <Printer class="w-4 h-4" />
              Print
            </button>
            <button
              v-if="invoice.status == 'pending'"
              @click="router.push(`/transactions/edit/${invoice.id}`)"
              class="px-4 py-2 bg-purple-100 hover:bg-purple-200 text-purple-700 rounded-xl text-sm font-medium transition-all flex items-center gap-2">
              <Edit2 class="w-4 h-4" />
              Edit
            </button>
          </div>
        </div>

        <!-- Invoice Info Grid -->
        <div :class="['grid grid-cols-1 gap-4', invoice.status != 'pending' ? 'md:grid-cols-4' : 'md:grid-cols-3']">
          <!-- Date -->
          <div class="p-4 bg-slate-50 rounded-xl border border-slate-200">
            <div class="flex items-center gap-2 text-slate-500 text-xs font-medium mb-1">
              <Calendar class="w-3.5 h-3.5" />
              Date
            </div>
            <p class="text-slate-800 font-bold">{{ formatDate(invoice.date) }}</p>
          </div>

          <!-- Status -->
          <div class="p-4 bg-slate-50 rounded-xl border border-slate-200">
            <p class="text-slate-500 text-xs font-medium mb-2">Status</p>
            <span :class="['inline-flex px-3 py-1.5 rounded-lg text-sm font-bold capitalize border-2', getStatusColor(invoice.status)]">
              {{ invoice.status }}
            </span>
          </div>

                    <!-- Payment Method -->
          <div v-if="invoice.status != 'pending'" class="p-4 bg-slate-50 rounded-xl border border-slate-200">
            <p class="text-slate-500 text-xs font-medium mb-2">Payment Method</p>
            <span class="inline-block px-3 py-1.5 bg-slate-100 text-slate-700 rounded-lg text-sm font-semibold capitalize">
              {{ invoice.payment_method }}
            </span>
          </div>

          <!-- Total Items -->
          <div class="p-4 bg-slate-50 rounded-xl border border-slate-200">
            <div class="flex items-center gap-2 text-slate-500 text-xs font-medium mb-1">
              <ShoppingCart class="w-3.5 h-3.5" />
              Total Items
            </div>
            <p class="text-slate-800 font-bold">{{ invoice.items.length }} {{ invoice.items.length === 1 ? 'item' : 'items' }}</p>
          </div>
        </div>
      </div>

      <!-- Transaction Items -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-xl font-bold text-slate-800 mb-5 flex items-center gap-2">
          <ShoppingCart class="w-5 h-5 text-purple-500" />
          Transaction Items
        </h2>

        <div class="space-y-4">
          <div
            v-for="(item, index) in invoice.items"
            :key="item.id"
            class="p-5 bg-gradient-to-br from-slate-50 to-purple-50 border-2 border-slate-200 rounded-xl">

            <div class="flex items-start gap-4">
              <!-- Item Number Badge -->
              <div :class="['w-12 h-12 rounded-xl bg-gradient-to-br flex items-center justify-center text-white font-bold shadow-md flex-shrink-0', getItemGradient(item.id)]">
                {{ index + 1 }}
              </div>

              <!-- Item Details -->
              <div class="flex-1 space-y-3">
                <!-- Service Name -->
                <div>
                  <div class="flex items-center gap-2 mb-1">
                    <Heart class="w-4 h-4 text-purple-500" />
                    <span class="text-xs text-slate-500 font-medium">Service</span>
                  </div>
                  <p class="text-lg font-bold text-slate-800">{{ services[item.service_id]?.name || 'Loading...' }}</p>
                </div>

                <!-- Customer & Therapist -->
                <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
                  <div class="p-3 bg-white rounded-lg border border-slate-200">
                    <div class="flex items-center gap-2 mb-1">
                      <Users class="w-3.5 h-3.5 text-slate-400" />
                      <span class="text-xs text-slate-500 font-medium">Customer</span>
                    </div>
                    <p class="font-semibold text-slate-800 text-sm">{{ customers[item.customer_id]?.name || 'Loading...' }}</p>
                    <p class="text-xs text-slate-500">{{ customers[item.customer_id]?.phone || '' }}</p>
                  </div>

                  <div class="p-3 bg-white rounded-lg border border-slate-200">
                    <div class="flex items-center gap-2 mb-1">
                      <UserCircle class="w-3.5 h-3.5 text-slate-400" />
                      <span class="text-xs text-slate-500 font-medium">Therapist</span>
                    </div>
                    <p class="font-semibold text-slate-800 text-sm">{{ therapists[item.therapist_id]?.name || 'Loading...' }}</p>
                    <!-- <p class="text-xs text-purple-600 font-bold">
                      Commission: {{ formatCurrency(calculateCommission(item.price, item.therapist_id)) }}
                      ({{ therapists[item.therapist_id]?.commission_rate || 0 }}%)
                    </p> -->
                  </div>
                </div>
              </div>

              <!-- Price -->
              <div class="text-right">
                <p class="text-xs text-slate-500 mb-1">Price</p>
                <p class="text-2xl font-bold text-emerald-600">{{ formatCurrency(item.price) }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Pricing Summary -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-xl font-bold text-slate-800 mb-5 flex items-center gap-2">
          <DollarSign class="w-5 h-5 text-purple-500" />
          Payment Summary
        </h2>

        <div class="space-y-4">
          <!-- Sub Total -->
          <div class="flex justify-between items-center pb-3">
            <span class="text-slate-600 font-medium">Sub Total</span>
            <span class="text-slate-800 font-bold text-xl">{{ formatCurrency(invoice.sub_total) }}</span>
          </div>

          <!-- Discount -->
          <div class="flex justify-between items-center pb-3">
            <span class="text-slate-600 font-medium">Discount</span>
            <span class="text-rose-600 font-bold text-xl">-{{ formatCurrency(invoice.discount_amount) }}</span>
          </div>

          <!-- Grand Total -->
          <div class="flex justify-between items-center pt-4 border-t-2 border-slate-300">
            <span class="text-slate-800 font-bold text-xl">Grand Total</span>
            <span class="text-emerald-600 font-bold text-3xl">{{ formatCurrency(invoice.grand_total) }}</span>
          </div>
        </div>

        <!-- Total Commission Summary -->
        <!-- <div class="mt-6 p-4 bg-purple-50 border-2 border-purple-200 rounded-xl">
          <div class="flex justify-between items-center">
            <span class="text-purple-700 font-medium">Total Therapist Commissions</span>
            <span class="text-purple-900 font-bold text-xl">
              {{ formatCurrency(invoice.items.reduce((sum, item) => sum + calculateCommission(item.price, item.therapist_id), 0)) }}
            </span>
          </div>
        </div> -->
      </div>

    </div>
  </div>
</template>

<style scoped>
/* Smooth transitions */
* {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Print styles */
@media print {
  .no-print {
    display: none !important;
  }
}

/* Smooth font rendering */
body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
