package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/4thel00z/shopify-graphql/graph/generated"
	"github.com/4thel00z/shopify-graphql/graph/model"
)

func (r *mutationResolver) AppCreditCreate(ctx context.Context, description string, amount model.MoneyInput, test *bool) (*model.AppCreditCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AppPurchaseOneTimeCreate(ctx context.Context, name string, price model.MoneyInput, returnURL string, test *bool) (*model.AppPurchaseOneTimeCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AppSubscriptionCancel(ctx context.Context, id string) (*model.AppSubscriptionCancelPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AppSubscriptionCreate(ctx context.Context, name string, lineItems []*model.AppSubscriptionLineItemInput, test *bool, trialDays *int, returnURL string) (*model.AppSubscriptionCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AppSubscriptionLineItemUpdate(ctx context.Context, id string, cappedAmount model.MoneyInput) (*model.AppSubscriptionLineItemUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AppUsageRecordCreate(ctx context.Context, subscriptionLineItemID string, price model.MoneyInput, description string) (*model.AppUsageRecordCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) BulkOperationCancel(ctx context.Context, id string) (*model.BulkOperationCancelPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) BulkOperationRunQuery(ctx context.Context, query string) (*model.BulkOperationRunQueryPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CollectionAddProducts(ctx context.Context, id string, productIds []string) (*model.CollectionAddProductsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CollectionCreate(ctx context.Context, input model.CollectionInput) (*model.CollectionCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CollectionDelete(ctx context.Context, input model.CollectionDeleteInput) (*model.CollectionDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CollectionPublish(ctx context.Context, input model.CollectionPublishInput) (*model.CollectionPublishPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CollectionRemoveProducts(ctx context.Context, id string, productIds []string) (*model.CollectionRemoveProductsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CollectionReorderProducts(ctx context.Context, id string, moves []*model.MoveInput) (*model.CollectionReorderProductsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CollectionUnpublish(ctx context.Context, input model.CollectionUnpublishInput) (*model.CollectionUnpublishPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CollectionUpdate(ctx context.Context, input model.CollectionInput) (*model.CollectionUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CustomerAddTaxExemptions(ctx context.Context, customerID string, taxExemptions []model.TaxExemption) (*model.CustomerAddTaxExemptionsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CustomerCreate(ctx context.Context, input model.CustomerInput) (*model.CustomerCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CustomerDelete(ctx context.Context, input model.CustomerDeleteInput) (*model.CustomerDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CustomerGenerateAccountActivationURL(ctx context.Context, customerID string) (*model.CustomerGenerateAccountActivationURLPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CustomerRemoveTaxExemptions(ctx context.Context, customerID string, taxExemptions []model.TaxExemption) (*model.CustomerRemoveTaxExemptionsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CustomerReplaceTaxExemptions(ctx context.Context, customerID string, taxExemptions []model.TaxExemption) (*model.CustomerReplaceTaxExemptionsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CustomerUpdate(ctx context.Context, input model.CustomerInput) (*model.CustomerUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CustomerUpdateDefaultAddress(ctx context.Context, customerID string, addressID string) (*model.CustomerUpdateDefaultAddressPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeliveryProfileCreate(ctx context.Context, profile model.DeliveryProfileInput) (*model.DeliveryProfileCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeliveryProfileRemove(ctx context.Context, id string) (*model.DeliveryProfileRemovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeliveryProfileUpdate(ctx context.Context, id string, profile model.DeliveryProfileInput, leaveLegacyModeProfiles *bool) (*model.DeliveryProfileUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeliverySettingUpdate(ctx context.Context, setting model.DeliverySettingInput) (*model.DeliverySettingUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeliveryShippingOriginAssign(ctx context.Context, locationID string) (*model.DeliveryShippingOriginAssignPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountAutomaticActivate(ctx context.Context, id string) (*model.DiscountAutomaticActivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountAutomaticBasicCreate(ctx context.Context, automaticBasicDiscount model.DiscountAutomaticBasicInput) (*model.DiscountAutomaticBasicCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountAutomaticBasicUpdate(ctx context.Context, id string, automaticBasicDiscount model.DiscountAutomaticBasicInput) (*model.DiscountAutomaticBasicUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountAutomaticBulkDelete(ctx context.Context, search *string, savedSearchID *string, ids []string) (*model.DiscountAutomaticBulkDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountAutomaticBxgyCreate(ctx context.Context, automaticBxgyDiscount model.DiscountAutomaticBxgyInput) (*model.DiscountAutomaticBxgyCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountAutomaticBxgyUpdate(ctx context.Context, id string, automaticBxgyDiscount model.DiscountAutomaticBxgyInput) (*model.DiscountAutomaticBxgyUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountAutomaticDeactivate(ctx context.Context, id string) (*model.DiscountAutomaticDeactivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountAutomaticDelete(ctx context.Context, id string) (*model.DiscountAutomaticDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeActivate(ctx context.Context, id string) (*model.DiscountCodeActivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeBasicCreate(ctx context.Context, basicCodeDiscount model.DiscountCodeBasicInput) (*model.DiscountCodeBasicCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeBasicUpdate(ctx context.Context, id string, basicCodeDiscount model.DiscountCodeBasicInput) (*model.DiscountCodeBasicUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeBulkActivate(ctx context.Context, search *string, savedSearchID *string, ids []string) (*model.DiscountCodeBulkActivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeBulkDeactivate(ctx context.Context, search *string, savedSearchID *string, ids []string) (*model.DiscountCodeBulkDeactivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeBulkDelete(ctx context.Context, search *string, savedSearchID *string, ids []string) (*model.DiscountCodeBulkDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeBxgyCreate(ctx context.Context, bxgyCodeDiscount model.DiscountCodeBxgyInput) (*model.DiscountCodeBxgyCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeBxgyUpdate(ctx context.Context, id string, bxgyCodeDiscount model.DiscountCodeBxgyInput) (*model.DiscountCodeBxgyUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeDeactivate(ctx context.Context, id string) (*model.DiscountCodeDeactivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeDelete(ctx context.Context, id string) (*model.DiscountCodeDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeFreeShippingCreate(ctx context.Context, freeShippingCodeDiscount model.DiscountCodeFreeShippingInput) (*model.DiscountCodeFreeShippingCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeFreeShippingUpdate(ctx context.Context, id string, freeShippingCodeDiscount model.DiscountCodeFreeShippingInput) (*model.DiscountCodeFreeShippingUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscountCodeRedeemCodeBulkDelete(ctx context.Context, discountID string, search *string, savedSearchID *string, ids []string) (*model.DiscountCodeRedeemCodeBulkDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DraftOrderCalculate(ctx context.Context, input model.DraftOrderInput) (*model.DraftOrderCalculatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DraftOrderComplete(ctx context.Context, id string, paymentPending *bool) (*model.DraftOrderCompletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DraftOrderCreate(ctx context.Context, input model.DraftOrderInput) (*model.DraftOrderCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DraftOrderDelete(ctx context.Context, input model.DraftOrderDeleteInput) (*model.DraftOrderDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DraftOrderInvoicePreview(ctx context.Context, id string, email *model.EmailInput) (*model.DraftOrderInvoicePreviewPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DraftOrderInvoiceSend(ctx context.Context, id string, email *model.EmailInput) (*model.DraftOrderInvoiceSendPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DraftOrderUpdate(ctx context.Context, id string, input model.DraftOrderInput) (*model.DraftOrderUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EventBridgeWebhookSubscriptionCreate(ctx context.Context, topic model.WebhookSubscriptionTopic, webhookSubscription model.EventBridgeWebhookSubscriptionInput) (*model.EventBridgeWebhookSubscriptionCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EventBridgeWebhookSubscriptionUpdate(ctx context.Context, id string, webhookSubscription model.EventBridgeWebhookSubscriptionInput) (*model.EventBridgeWebhookSubscriptionUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FlowTriggerReceive(ctx context.Context, body string) (*model.FlowTriggerReceivePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentCancel(ctx context.Context, id string) (*model.FulfillmentCancelPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentCreate(ctx context.Context, input model.FulfillmentInput) (*model.FulfillmentCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentCreateV2(ctx context.Context, fulfillment model.FulfillmentV2Input, message *string) (*model.FulfillmentCreateV2Payload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderAcceptCancellationRequest(ctx context.Context, id string, message *string) (*model.FulfillmentOrderAcceptCancellationRequestPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderAcceptFulfillmentRequest(ctx context.Context, id string, message *string) (*model.FulfillmentOrderAcceptFulfillmentRequestPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderCancel(ctx context.Context, id string) (*model.FulfillmentOrderCancelPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderClose(ctx context.Context, id string, message *string) (*model.FulfillmentOrderClosePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderMove(ctx context.Context, id string, newLocationID string) (*model.FulfillmentOrderMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderRejectCancellationRequest(ctx context.Context, id string, message *string) (*model.FulfillmentOrderRejectCancellationRequestPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderRejectFulfillmentRequest(ctx context.Context, id string, message *string) (*model.FulfillmentOrderRejectFulfillmentRequestPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderSubmitCancellationRequest(ctx context.Context, id string, message *string) (*model.FulfillmentOrderSubmitCancellationRequestPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentOrderSubmitFulfillmentRequest(ctx context.Context, id string, message *string, notifyCustomer *bool, fulfillmentOrderLineItems []*model.FulfillmentOrderLineItemInput, shippingMethod *string) (*model.FulfillmentOrderSubmitFulfillmentRequestPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentServiceCreate(ctx context.Context, name string, callbackURL *string, trackingSupport *bool) (*model.FulfillmentServiceCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentServiceDelete(ctx context.Context, id string, destinationLocationID *string) (*model.FulfillmentServiceDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentServiceUpdate(ctx context.Context, id string, name *string, callbackURL *string, trackingSupport *bool, fulfillmentOrdersOptIn *bool) (*model.FulfillmentServiceUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentTrackingInfoUpdate(ctx context.Context, fulfillmentID string, trackingInfoUpdateInput model.TrackingInfoUpdateInput) (*model.FulfillmentTrackingInfoUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FulfillmentTrackingInfoUpdateV2(ctx context.Context, fulfillmentID string, trackingInfoInput model.FulfillmentTrackingInput, notifyCustomer *bool) (*model.FulfillmentTrackingInfoUpdateV2Payload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InventoryActivate(ctx context.Context, inventoryItemID string, locationID string, available *int) (*model.InventoryActivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InventoryAdjustQuantity(ctx context.Context, input model.InventoryAdjustQuantityInput) (*model.InventoryAdjustQuantityPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InventoryBulkAdjustQuantityAtLocation(ctx context.Context, inventoryItemAdjustments []*model.InventoryAdjustItemInput, locationID string) (*model.InventoryBulkAdjustQuantityAtLocationPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InventoryDeactivate(ctx context.Context, inventoryLevelID string) (*model.InventoryDeactivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InventoryItemUpdate(ctx context.Context, id string, input model.InventoryItemUpdateInput) (*model.InventoryItemUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) KitSkillTriggerRequest(ctx context.Context, id string, locale model.KitSkillLocale, placeholders *string) (*model.KitSkillTriggerRequestPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MarketingActivityCreate(ctx context.Context, input model.MarketingActivityCreateInput) (*model.MarketingActivityCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MarketingActivityUpdate(ctx context.Context, input model.MarketingActivityUpdateInput) (*model.MarketingActivityUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MarketingEngagementCreate(ctx context.Context, marketingActivityID string, marketingEngagement model.MarketingEngagementInput) (*model.MarketingEngagementCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MetafieldDelete(ctx context.Context, input model.MetafieldDeleteInput) (*model.MetafieldDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MetafieldStorefrontVisibilityCreate(ctx context.Context, input model.MetafieldStorefrontVisibilityInput) (*model.MetafieldStorefrontVisibilityCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MetafieldStorefrontVisibilityDelete(ctx context.Context, id string) (*model.MetafieldStorefrontVisibilityDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderCapture(ctx context.Context, input model.OrderCaptureInput) (*model.OrderCapturePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderClose(ctx context.Context, input model.OrderCloseInput) (*model.OrderClosePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderEditAddCustomItem(ctx context.Context, id string, title string, locationID *string, price model.MoneyInput, quantity int, taxable *bool, requiresShipping *bool) (*model.OrderEditAddCustomItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderEditAddLineItemDiscount(ctx context.Context, id string, lineItemID string, discount model.OrderEditAppliedDiscountInput) (*model.OrderEditAddLineItemDiscountPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderEditAddVariant(ctx context.Context, id string, variantID string, locationID *string, quantity int, allowDuplicates *bool) (*model.OrderEditAddVariantPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderEditBegin(ctx context.Context, id string) (*model.OrderEditBeginPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderEditCommit(ctx context.Context, id string, notifyCustomer *bool, staffNote *string) (*model.OrderEditCommitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderEditRemoveLineItemDiscount(ctx context.Context, id string, discountApplicationID string) (*model.OrderEditRemoveLineItemDiscountPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderEditSetQuantity(ctx context.Context, id string, lineItemID string, quantity int, restock *bool, locationID *string) (*model.OrderEditSetQuantityPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderMarkAsPaid(ctx context.Context, input model.OrderMarkAsPaidInput) (*model.OrderMarkAsPaidPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderOpen(ctx context.Context, input model.OrderOpenInput) (*model.OrderOpenPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderUpdate(ctx context.Context, input model.OrderInput) (*model.OrderUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PriceRuleActivate(ctx context.Context, id string) (*model.PriceRuleActivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PriceRuleCreate(ctx context.Context, priceRule model.PriceRuleInput, priceRuleDiscountCode *model.PriceRuleDiscountCodeInput) (*model.PriceRuleCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PriceRuleDeactivate(ctx context.Context, id string) (*model.PriceRuleDeactivatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PriceRuleDelete(ctx context.Context, id string) (*model.PriceRuleDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PriceRuleDiscountCodeCreate(ctx context.Context, priceRuleID string, code string) (*model.PriceRuleDiscountCodeCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PriceRuleDiscountCodeUpdate(ctx context.Context, priceRuleID string, code string) (*model.PriceRuleDiscountCodeUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PriceRuleUpdate(ctx context.Context, id string, priceRule model.PriceRuleInput, priceRuleDiscountCode *model.PriceRuleDiscountCodeInput) (*model.PriceRuleUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PrivateMetafieldDelete(ctx context.Context, input model.PrivateMetafieldDeleteInput) (*model.PrivateMetafieldDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PrivateMetafieldUpsert(ctx context.Context, input model.PrivateMetafieldInput) (*model.PrivateMetafieldUpsertPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductAppendImages(ctx context.Context, input model.ProductAppendImagesInput) (*model.ProductAppendImagesPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductChangeStatus(ctx context.Context, productID string, status model.ProductStatus) (*model.ProductChangeStatusPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductCreate(ctx context.Context, input model.ProductInput, media []*model.CreateMediaInput) (*model.ProductCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductCreateMedia(ctx context.Context, productID string, media []*model.CreateMediaInput) (*model.ProductCreateMediaPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductDelete(ctx context.Context, input model.ProductDeleteInput) (*model.ProductDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductDeleteImages(ctx context.Context, id string, imageIds []string) (*model.ProductDeleteImagesPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductDeleteMedia(ctx context.Context, productID string, mediaIds []string) (*model.ProductDeleteMediaPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductDuplicate(ctx context.Context, productID string, newTitle string, newStatus *model.ProductStatus, includeImages *bool) (*model.ProductDuplicatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductImageUpdate(ctx context.Context, productID string, image model.ImageInput) (*model.ProductImageUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductPublish(ctx context.Context, input model.ProductPublishInput) (*model.ProductPublishPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductReorderImages(ctx context.Context, id string, moves []*model.MoveInput) (*model.ProductReorderImagesPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductReorderMedia(ctx context.Context, id string, moves []*model.MoveInput) (*model.ProductReorderMediaPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductUnpublish(ctx context.Context, input model.ProductUnpublishInput) (*model.ProductUnpublishPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductUpdate(ctx context.Context, input model.ProductInput) (*model.ProductUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductUpdateMedia(ctx context.Context, productID string, media []*model.UpdateMediaInput) (*model.ProductUpdateMediaPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductVariantAppendMedia(ctx context.Context, productID string, variantMedia []*model.ProductVariantAppendMediaInput) (*model.ProductVariantAppendMediaPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductVariantCreate(ctx context.Context, input model.ProductVariantInput) (*model.ProductVariantCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductVariantDelete(ctx context.Context, id string) (*model.ProductVariantDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductVariantDetachMedia(ctx context.Context, productID string, variantMedia []*model.ProductVariantDetachMediaInput) (*model.ProductVariantDetachMediaPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductVariantUpdate(ctx context.Context, input model.ProductVariantInput) (*model.ProductVariantUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PublishablePublish(ctx context.Context, id string, input []*model.PublicationInput) (*model.PublishablePublishPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PublishablePublishToCurrentChannel(ctx context.Context, id string) (*model.PublishablePublishToCurrentChannelPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PublishableUnpublish(ctx context.Context, id string, input []*model.PublicationInput) (*model.PublishableUnpublishPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PublishableUnpublishToCurrentChannel(ctx context.Context, id string) (*model.PublishableUnpublishToCurrentChannelPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefundCreate(ctx context.Context, input model.RefundInput) (*model.RefundCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SavedSearchCreate(ctx context.Context, input model.SavedSearchCreateInput) (*model.SavedSearchCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SavedSearchDelete(ctx context.Context, input model.SavedSearchDeleteInput) (*model.SavedSearchDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SavedSearchUpdate(ctx context.Context, input model.SavedSearchUpdateInput) (*model.SavedSearchUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ScriptTagCreate(ctx context.Context, input model.ScriptTagInput) (*model.ScriptTagCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ScriptTagDelete(ctx context.Context, id string) (*model.ScriptTagDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ScriptTagUpdate(ctx context.Context, id string, input model.ScriptTagInput) (*model.ScriptTagUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ShippingPackageDelete(ctx context.Context, id string) (*model.ShippingPackageDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ShippingPackageMakeDefault(ctx context.Context, id string) (*model.ShippingPackageMakeDefaultPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ShippingPackageUpdate(ctx context.Context, id string) (*model.ShippingPackageUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ShopLocaleDisable(ctx context.Context, locale string) (*model.ShopLocaleDisablePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ShopLocaleEnable(ctx context.Context, locale string) (*model.ShopLocaleEnablePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ShopLocaleUpdate(ctx context.Context, locale string, shopLocale model.ShopLocaleInput) (*model.ShopLocaleUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ShopPolicyUpdate(ctx context.Context, shopPolicy model.ShopPolicyInput) (*model.ShopPolicyUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StagedUploadTargetGenerate(ctx context.Context, input model.StagedUploadTargetGenerateInput) (*model.StagedUploadTargetGeneratePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StagedUploadTargetsGenerate(ctx context.Context, input []*model.StageImageInput) (*model.StagedUploadTargetsGeneratePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StagedUploadsCreate(ctx context.Context, input []*model.StagedUploadInput) (*model.StagedUploadsCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StorefrontAccessTokenCreate(ctx context.Context, input model.StorefrontAccessTokenInput) (*model.StorefrontAccessTokenCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StorefrontAccessTokenDelete(ctx context.Context, input model.StorefrontAccessTokenDeleteInput) (*model.StorefrontAccessTokenDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) TagsAdd(ctx context.Context, id string, tags []string) (*model.TagsAddPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) TagsRemove(ctx context.Context, id string, tags []string) (*model.TagsRemovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) TranslationsRegister(ctx context.Context, resourceID string, translations []*model.TranslationInput) (*model.TranslationsRegisterPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) TranslationsRemove(ctx context.Context, resourceID string, translationKeys []string, locales []string) (*model.TranslationsRemovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) WebhookSubscriptionCreate(ctx context.Context, topic model.WebhookSubscriptionTopic, webhookSubscription model.WebhookSubscriptionInput) (*model.WebhookSubscriptionCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) WebhookSubscriptionDelete(ctx context.Context, id string) (*model.WebhookSubscriptionDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) WebhookSubscriptionUpdate(ctx context.Context, id string, webhookSubscription model.WebhookSubscriptionInput) (*model.WebhookSubscriptionUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
