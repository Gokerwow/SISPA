<script setup lang="ts">
import type { Customer } from '@/types/customer';
import { formatForInput } from '@/utils/date';
import { reactive, watch } from 'vue';
import {
  User,
  Phone,
  MapPin,
  Calendar,
  Save,
  X,
  Gift,
} from 'lucide-vue-next'

const form = reactive<Customer>({
  name: '',
  phone: '',
  birth_date: '',
  gender: '',
  address: '',
  note: ''
})

const props = defineProps<{
  initialData?: Customer,
  isSubmitting?: boolean
}>()

const emits = defineEmits<{
  (e: 'submit', data: Customer): void;
}>()

// Use a watcher to catch the data whenever it arrives
watch(() => props.initialData, (newData) => {
  if (newData) {
    Object.assign(form, newData)
    // 2. Handle the "Zero Date" from Go
    if (newData.birth_date && !newData.birth_date.startsWith('0001')) {
      form.birth_date = formatForInput(newData.birth_date)
    } else {
      form.birth_date = '' // Keep it empty so the input shows "dd/mm/yyyy"
    }
  }
}, { immediate: true }) // 'immediate' makes it run even if data is already there on mount

const handleSubmit = () => {
  console.log(form)
  emits('submit', { ...form })
}


</script>

<template>
  <div class=" space-y-6">

    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1
          class="text-3xl font-bold bg-gradient-to-r from-rose-600 via-purple-600 to-indigo-600 bg-clip-text text-transparent mb-2">
          Add New Customer
        </h1>
        <p class="text-slate-500 text-sm">Create a new customer profile for your spa</p>
      </div>
      <button @click="$router.back()" class="p-3 hover:bg-slate-100 rounded-xl transition-all text-slate-600">
        <X class="w-5 h-5" />
      </button>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">

      <!-- Personal Information -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <User class="w-5 h-5 text-purple-500" />
          Personal Information
        </h2>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          <!-- Full Name -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Full Name <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <User class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
              <input v-model="form.name" type="text" required placeholder="John Doe"
                class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all" />
            </div>
          </div>

          <!-- Phone -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Phone Number <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <Phone class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
              <input v-model="form.phone" type="tel" required placeholder="+62 (555) 000-0000"
                class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all" />
            </div>
          </div>

          <!-- Date of Birth -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Date of Birth
            </label>
            <div class="relative">
              <Calendar class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
              <input v-model="form.birth_date" type="date" required
                class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all" />
            </div>
          </div>

          <!-- Gender -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Gender
            </label>
            <select v-model="form.gender" required
              class="w-full px-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all">
              <option value="">Select gender</option>
              <option value="female">Female</option>
              <option value="male">Male</option>
            </select>
          </div>
        </div>
      </div>
      <!-- Address Information -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <MapPin class="w-5 h-5 text-purple-500" />
          Address
        </h2>

        <div>
          <label class="block text-sm font-medium text-slate-700 mb-2">
            Full Address
          </label>
          <div class="relative">
            <MapPin class="absolute left-4 top-4 w-5 h-5 text-slate-400" />
            <textarea v-model="form.address" rows="3" required
              placeholder="123 Main Street, Apt 4B, New York, NY 10001, United States"
              class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all resize-none"></textarea>
          </div>
        </div>
      </div>

      <!-- Additional Notes -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-4 flex items-center gap-2">
          <Gift class="w-5 h-5 text-purple-500" />
          Additional Notes
        </h2>

        <textarea v-model="form.note" rows="4"
          placeholder="Any allergies, special requests, or important information about this customer..."
          class="w-full px-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all resize-none"></textarea>
      </div>

      <!-- Action Buttons -->
      <div class="flex items-center justify-end gap-4">
        <button type="button" @click="$router.back()"
          class="px-6 py-3 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-2xl text-sm font-medium transition-all flex items-center gap-2">
          <X class="w-4 h-4" />
          Cancel
        </button>

        <button type="submit" :disabled="isSubmitting"
          class="px-8 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl hover:shadow-purple-300/50 transition-all flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
          <Save class="w-4 h-4" />
          <span v-if="!isSubmitting">Save Customer</span>
          <span v-else>Saving...</span>
        </button>
      </div>

    </form>
  </div>
</template>
