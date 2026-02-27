import type { TransactionRequest } from "@/types/transaction"
import { useToast } from "vue-toastification"

// Validate form
const toast = useToast()

export const validateForm = (formData: TransactionRequest) => {
  if (!formData.date) {
    toast.error('Please select a date')
    return false
  }

  if (formData.grand_total <= 0) {
    toast.error('Grand total cannot be minus or 0, please check the discount and prices of the services')
    return false
  }

  if (formData.items.length === 0) {
    toast.error('Please add at least one transaction item')
    return false
  }

  for (let i = 0; i < formData.items.length; i++) {
    const item = formData.items[i]
    if (!item?.customer_id || !item.therapist_id || !item.service_id) {
      toast.error(`Please fill all fields for item ${i + 1}`)
      return false
    }
  }
  return true
}
