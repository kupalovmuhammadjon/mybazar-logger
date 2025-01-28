package logger

// Validation Error Codes
const (
	// 1001: Required field missing in the API request.
	ErrReqFieldMissing = 1001
	// 1002: Invalid data format or type provided.
	ErrInvalidData = 1002
	// 1003: Value exceeds allowed range.
	ErrValueExceedsRange = 1003
	// 1004: Unsupported file type uploaded.
	ErrUnsupportedFile = 1004
	// 1005: Duplicate data found in the request.
	ErrDuplicateData = 1005
	// 1006: Invalid query parameter or filter provided.
	ErrInvalidQuery = 1006
)

// Authentication Error Codes
const (
	// 2001: User is not authenticated.
	ErrNotAuthenticated = 2001
	// 2002: User does not have permission to access this resource.
	ErrPermissionDenied = 2002
	// 2003: Invalid or expired authentication token.
	ErrInvalidToken = 2003
	// 2004: Account temporarily locked due to multiple failed attempts.
	ErrAccountLocked = 2004
	// 2005: Session expired; re-authentication required.
	ErrSessionExpired = 2005
)

// Resource Error Codes
const (
	// 3001: Resource not found (e.g., product, order).
	ErrResourceNotFound = 3001
	// 3002: Resource is currently unavailable or locked.
	ErrResourceLocked = 3002
	// 3003: Insufficient inventory for requested product.
	ErrInsufficientInventory = 3003
	// 3004: Resource has been archived or deleted.
	ErrResourceArchived = 3004
	// 3005: Dependency not found (e.g., related resource missing).
	ErrDependencyNotFound = 3005
)

// System Error Codes
const (
	// 4001: Internal server error.
	ErrInternalServer = 4001
	// 4002: Service is temporarily unavailable.
	ErrServiceUnavailable = 4002
	// 4003: Database connection error.
	ErrDatabaseError = 4003
	// 4004: Cache synchronization failed.
	ErrCacheSyncFailed = 4004
	// 4005: Unexpected behavior in background job processing.
	ErrJobProcessingError = 4005
)

// Integration Error Codes
const (
	// 5001: Third-party API returned an error.
	ErrAPIError = 5001
	// 5002: Failed to connect to an external service.
	ErrConnectionFailed = 5002
	// 5003: Timeout while waiting for a third-party API response.
	ErrAPITimeout = 5003
	// 5004: Invalid response received from third-party service.
	ErrInvalidAPIResponse = 5004
	// 5005: API quota limit reached for external service.
	ErrAPILimitReached = 5005
)

// Business Logic Error Codes
const (
	// 6001: Order cannot be processed due to invalid status.
	ErrInvalidOrderStatus = 6001
	// 6002: Merchant quota exceeded for daily requests.
	ErrMerchantQuotaExceeded = 6002
	// 6003: Payment gateway rejected the transaction.
	ErrPaymentRejected = 6003
	// 6004: Refund cannot be processed due to insufficient balance.
	ErrRefundFailed = 6004
	// 6005: Promotion code is invalid or expired.
	ErrInvalidPromoCode = 6005
	// 6006: Order cancellation window has passed.
	ErrCancellationWindowClosed = 6006
)
