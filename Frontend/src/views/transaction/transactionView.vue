<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { FileText, Plus, Search, DollarSign, Edit2, Eye, Filter, TrendingUp, CreditCard, Calendar } from 'lucide-vue-next'
import { RouterLink } from 'vue-router'
import { GetAllHeader, updateStatusTransaction } from '@/services/transactions'
import { useToast } from 'vue-toastification'
import type { Invoice, paymentMethod, UpdateStatusRequest } from '@/types/transaction'

const invoices = ref<Invoice[]>([])
const loading = ref(true)
const error = ref("")
const searchQuery = ref("")
const refreshTrigger = ref(0)
const toast = useToast()

// Fetch Logic
const fetchInvoices = async () => {
  try {
    const response = await GetAllHeader()
    invoices.value = response.data
    console.log(invoices.value)
  } catch (err) {
    error.value = (err as Error).message
    console.log(error.value)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchInvoices()
})

watch(refreshTrigger, () => {
  fetchInvoices()
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
    month: 'short',
    day: 'numeric'
  })
}

// Get status badge color
const getStatusColor = (status: string) => {
  const colors: Record<string, string> = {
    'paid': 'bg-emerald-100 text-emerald-700',
    'pending': 'bg-amber-100 text-amber-700',
    'cancelled': 'bg-red-100 text-red-700'
  }
  return colors[status?.toLowerCase()] || 'bg-slate-100 text-slate-700'
}

// Get invoice icon gradient
const getInvoiceGradient = (id: number) => {
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

const updateStatus = async (invoice: Invoice, newStatus: 'pending' | 'paid' | 'cancelled', paymentMethod?: paymentMethod) => {
  try {
    const payload : UpdateStatusRequest = {
      status: newStatus,
      payment_method: paymentMethod ?? null
    }
    const response = await updateStatusTransaction(invoice.id, payload)
    // Update local data
    invoice.status = newStatus
    refreshTrigger.value += 1
    console.log('Status updated', response)
    toast.success(`Status updated for Invoice ${invoice.nota_number}`)
  } catch (error) {
    console.error('Error updating status:', error)
    toast.error(`Error updating status: ${(error as Error).message || 'Unknown error'}`);
  }
}
// Calculate stats
const totalRevenue = ref(0)
const paidInvoices = ref(0)
const pendingAmount = ref(0)

watch(invoices, (newInvoices) => {
  totalRevenue.value = newInvoices
    .filter(i => i.status === 'paid')
    .reduce((sum, i) => sum + i.grand_total, 0)

  paidInvoices.value = newInvoices.filter(i => i.status === 'paid').length

  pendingAmount.value = newInvoices
    .filter(i => i.status === 'pending')
    .reduce((sum, i) => sum + i.grand_total, 0)
}, { immediate: true })
</script>

<template>
  <div class="space-y-6">
    <!-- Header Section -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1
          class="text-3xl font-bold bg-gradient-to-r from-rose-600 via-purple-600 to-indigo-600 bg-clip-text text-transparent mb-2">
          Transactions
        </h1>
        <p class="text-slate-500 text-sm">Manage invoices and payment records</p>
      </div>
      <RouterLink to="/transactions/create"
        class="px-6 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl hover:shadow-purple-300/50 transition-all flex items-center gap-2">
        <Plus class="w-4 h-4" />
        <span>New Transaction</span>
      </RouterLink>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div
            class="w-12 h-12 rounded-xl bg-gradient-to-br from-emerald-400 to-teal-500 flex items-center justify-center shadow-lg shadow-emerald-300/30">
            <TrendingUp class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">Total Revenue</p>
            <p class="text-2xl font-bold text-slate-800">{{ formatCurrency(totalRevenue) }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div
            class="w-12 h-12 rounded-xl bg-gradient-to-br from-blue-400 to-indigo-500 flex items-center justify-center shadow-lg shadow-blue-300/30">
            <CreditCard class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">Paid Invoices</p>
            <p class="text-2xl font-bold text-slate-800">{{ paidInvoices }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div
            class="w-12 h-12 rounded-xl bg-gradient-to-br from-amber-400 to-orange-500 flex items-center justify-center shadow-lg shadow-amber-300/30">
            <FileText class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">Pending Amount</p>
            <p class="text-2xl font-bold text-slate-800">{{ formatCurrency(pendingAmount) }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Search & Filter Bar -->
    <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-4 border border-white/60 shadow-lg">
      <div class="flex flex-col sm:flex-row gap-3">
        <div class="flex-1 relative">
          <Search class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
          <input v-model="searchQuery" type="search" placeholder="Search by invoice number..."
            class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all" />
        </div>
        <button
          class="px-5 py-3 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-xl text-sm font-medium transition-all flex items-center justify-center gap-2">
          <Filter class="w-4 h-4" />
          <span>Filters</span>
        </button>
      </div>
    </div>

    <!-- Error Message -->
    <div v-if="error" class="bg-red-50 border-l-4 border-red-500 rounded-xl p-5 shadow-lg">
      <div class="flex items-start gap-3">
        <div class="w-10 h-10 rounded-xl bg-red-100 flex items-center justify-center flex-shrink-0">
          <span class="text-red-600 text-xl">⚠️</span>
        </div>
        <div>
          <p class="font-bold text-red-800 mb-1">Connection Error</p>
          <p class="text-red-700 text-sm">{{ error }}</p>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="bg-white/80 backdrop-blur-xl rounded-2xl border border-white/60 shadow-lg p-20">
      <div class="text-center">
        <div
          class="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gradient-to-r from-rose-400 via-purple-400 to-indigo-400 mb-4 animate-pulse">
          <FileText class="w-8 h-8 text-white" />
        </div>
        <p class="text-slate-500 font-medium">Loading transactions...</p>
      </div>
    </div>

    <!-- Invoices List -->
    <div v-else class="bg-white/80 backdrop-blur-xl rounded-2xl border border-white/60 shadow-lg">
      <!-- Table for larger screens -->
      <div class="hidden lg:block">
        <table class="w-full">
          <thead>
            <tr class="border-b border-slate-100 bg-slate-50/50">
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Invoice</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Date</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Sub Total</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Discount</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Grand Total</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Status</th>
              <th class="text-center p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Payment Method</th>
              <th class="text-right p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100">
            <tr v-for="invoice in invoices" :key="invoice.id" class="hover:bg-slate-50/50 transition-colors group">
              <td class="p-5">
                <div class="flex items-center gap-3">
                  <div
                    :class="['w-12 h-12 rounded-xl bg-gradient-to-br flex items-center justify-center shadow-md', getInvoiceGradient(invoice.id)]">
                    <FileText class="w-6 h-6 text-white" />
                  </div>
                  <div>
                    <p class="font-semibold text-slate-800">{{ invoice.nota_number }}</p>
                    <p class="text-xs text-slate-500">ID: #{{ invoice.id }}</p>
                  </div>
                </div>
              </td>
              <td class="p-5">
                <div class="flex items-center gap-2 text-slate-600">
                  <Calendar class="w-4 h-4 text-slate-400" />
                  <span class="text-sm">{{ formatDate(invoice.date) }}</span>
                </div>
              </td>
              <td class="p-5">
                <span class="text-sm font-semibold text-slate-700">{{ formatCurrency(invoice.sub_total) }}</span>
              </td>
              <td class="p-5">
                <span class="text-sm font-semibold text-rose-600">-{{ formatCurrency(invoice.discount_amount) }}</span>
              </td>
              <td class="p-5">
                <div class="flex items-center gap-1">
                  <DollarSign class="w-4 h-4 text-emerald-400" />
                  <span class="text-sm font-bold text-slate-800">{{ formatCurrency(invoice.grand_total) }}</span>
                </div>
              </td>
              <td class="p-5">
                <div class="relative group/status">
                  <button type="button"
                    :class="['px-3 py-1 rounded-lg text-xs font-bold capitalize cursor-pointer hover:shadow-md transition-all', getStatusColor(invoice.status)]">
                    {{ invoice.status }}
                  </button>

                  <!-- Status Dropdown -->
                  <div v-if="invoice.status === 'pending'"
                    class="absolute top-full right-0 mt-1 bg-white rounded-xl shadow-xl border border-slate-200 p-2 min-w-[140px] opacity-0 invisible group-hover/status:opacity-100 group-hover/status:visible transition-all duration-200 z-10">

                    <button @click="updateStatus(invoice, 'paid', 'cash')" type="button"
                      class="w-full px-3 py-2 text-left text-xs font-bold capitalize rounded-lg hover:bg-emerald-50 text-emerald-700 transition-colors flex items-center gap-2">
                      <span class="w-2 h-2 rounded-full bg-emerald-400"></span>
                      Paid (Cash)
                    </button>

                    <button @click="updateStatus(invoice, 'paid', 'qris')" type="button"
                      class="w-full px-3 py-2 text-left text-xs font-bold capitalize rounded-lg hover:bg-blue-50 text-blue-700 transition-colors flex items-center gap-2">
                      <span class="w-2 h-2 rounded-full bg-blue-400"></span>
                      Paid (QRIS)
                    </button>

                    <button @click="updateStatus(invoice, 'cancelled')" type="button"
                      class="w-full px-3 py-2 text-left text-xs font-bold capitalize rounded-lg hover:bg-red-50 text-red-700 transition-colors flex items-center gap-2">
                      <span class="w-2 h-2 rounded-full bg-red-400"></span>
                      Cancelled
                    </button>
                  </div>
                </div>
              </td>
              <td class="p-5 text-center">
                <span class="px-3 py-1 bg-slate-100 text-slate-700 rounded-lg text-xs font-semibold capitalize">
                  {{ invoice.payment_method ?? "-" }}
                </span>
              </td>
              <td class="p-5">
                <div class="flex items-center justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                  <button @click="$router.push(`/transactions/${invoice.id}`)"
                    class="p-2 hover:bg-blue-50 text-blue-600 rounded-lg transition-colors" title="View">
                    <Eye class="w-4 h-4" />
                  </button>
                  <button v-if="invoice.status == 'pending'" @click="$router.push(`/transactions/edit/${invoice.id}`)"
                    class="p-2 hover:bg-purple-50 text-purple-600 rounded-lg transition-colors" title="Edit">
                    <Edit2 class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Card view for mobile -->
      <div class="lg:hidden divide-y divide-slate-100">
        <div v-for="invoice in invoices" :key="invoice.id" class="p-5 hover:bg-slate-50/50 transition-colors">
          <div class="flex items-start gap-4">
            <div
              :class="['w-14 h-14 rounded-xl bg-gradient-to-br flex items-center justify-center shadow-md flex-shrink-0', getInvoiceGradient(invoice.id)]">
              <FileText class="w-7 h-7 text-white" />
            </div>
            <div class="flex-1 min-w-0">
              <div class="flex items-start justify-between mb-2">
                <div>
                  <h3 class="font-semibold text-slate-800 mb-1">{{ invoice.nota_number }}</h3>
                  <div class="flex items-center gap-2 text-xs text-slate-500 mb-2">
                    <Calendar class="w-3 h-3" />
                    <span>{{ formatDate(invoice.date) }}</span>
                  </div>
                </div>
                <div class="relative group/status">
                  <button type="button"
                    :class="['px-2 py-1 rounded-lg text-xs font-bold capitalize cursor-pointer hover:shadow-md transition-all', getStatusColor(invoice.status)]">
                    {{ invoice.status }}
                  </button>

                  <!-- Status Dropdown -->
                  <div v-if="invoice.status === 'pending'"
                    class="absolute top-full right-0 mt-1 bg-white rounded-xl shadow-xl border border-slate-200 p-2 min-w-[140px] opacity-0 invisible group-hover/status:opacity-100 group-hover/status:visible transition-all duration-200 z-10">

                    <button @click="updateStatus(invoice, 'paid')" type="button"
                      class="w-full px-3 py-2 text-left text-xs font-bold capitalize rounded-lg hover:bg-emerald-50 text-emerald-700 transition-colors flex items-center gap-2">
                      <span class="w-2 h-2 rounded-full bg-emerald-400"></span>
                      Paid
                    </button>

                    <button @click="updateStatus(invoice, 'cancelled')" type="button"
                      class="w-full px-3 py-2 text-left text-xs font-bold capitalize rounded-lg hover:bg-red-50 text-red-700 transition-colors flex items-center gap-2">
                      <span class="w-2 h-2 rounded-full bg-red-400"></span>
                      Cancelled
                    </button>
                  </div>
                </div>
              </div>
              <div class="space-y-2 mb-3 text-sm bg-slate-50 rounded-lg p-3">
                <div class="flex justify-between items-center">
                  <span class="text-slate-600">Sub Total:</span>
                  <span class="font-semibold text-slate-800">{{ formatCurrency(invoice.sub_total) }}</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-slate-600">Discount:</span>
                  <span class="font-semibold text-rose-600">-{{ formatCurrency(invoice.discount_amount) }}</span>
                </div>
                <div class="flex justify-between items-center pt-2 border-t border-slate-200">
                  <span class="text-slate-800 font-bold">Grand Total:</span>
                  <span class="font-bold text-emerald-600 text-lg">{{ formatCurrency(invoice.grand_total) }}</span>
                </div>
              </div>
              <div class="flex gap-2">
                <button @click="$router.push(`/transactions/${invoice.id}`)"
                  class="px-3 py-1.5 bg-blue-50 hover:bg-blue-100 text-blue-600 rounded-lg text-xs font-medium transition-colors flex items-center gap-1">
                  <Eye class="w-3 h-3" />
                  View
                </button>
                <button @click="$router.push(`/transactions/edit/${invoice.id}`)"
                  class="px-3 py-1.5 bg-purple-50 hover:bg-purple-100 text-purple-600 rounded-lg text-xs font-medium transition-colors flex items-center gap-1">
                  <Edit2 class="w-3 h-3" />
                  Edit
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="invoices.length === 0" class="p-20 text-center">
        <div
          class="inline-flex items-center justify-center w-20 h-20 rounded-full bg-gradient-to-r from-slate-100 to-slate-200 mb-4">
          <FileText class="w-10 h-10 text-slate-400" />
        </div>
        <h3 class="text-lg font-semibold text-slate-700 mb-2">No transactions yet</h3>
        <p class="text-slate-500 text-sm mb-6">Start recording your first transaction</p>
        <RouterLink to="/transactions/create"
          class="px-6 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl transition-all inline-flex items-center gap-2">
          <Plus class="w-4 h-4" />
          <span>Create First Transaction</span>
        </RouterLink>
      </div>
    </div>
  </div>
</template>


<style scoped>
/* Smooth transitions */
* {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Smooth font rendering */
body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
