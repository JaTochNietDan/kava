<!--
order: 3
-->

# Messages

Users claim rewards using `MsgClaimUSDXMintingReward` and `MsgClaimHardLiquidityProviderReward` messages.

```go
// MsgClaimUSDXMintingReward message type used to claim USDX minting rewards
type MsgClaimUSDXMintingReward struct {
  Sender         sdk.AccAddress `json:"sender" yaml:"sender"`
  MultiplierName string         `json:"multiplier_name" yaml:"multiplier_name"`
}

// MsgClaimHardLiquidityProviderReward message type used to claim Hard liquidity provider rewards
type MsgClaimHardLiquidityProviderReward struct {
  Sender         sdk.AccAddress `json:"sender" yaml:"sender"`
  MultiplierName string         `json:"multiplier_name" yaml:"multiplier_name"`
}
```

## State Modifications

* Accumulated rewards for active claims are transferred from the `kavadist` module account to the users account as vesting coins
* The number of coins transferred is determined by the multiplier in the message. For example, the multiplier equals 1.0, 100% of the claim's reward value is transferred. If the multiplier equals 0.5, 50% of the claim's reward value is transferred.
* The corresponding claim object is reset to zero in the store
