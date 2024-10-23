package sebp

import (
	"encoding/xml"

	"github.com/Merith-TK/se-workshop/shared"
)

type Metadata struct {
	XMLName        xml.Name `xml:"Definitions"`
	Text           string   `xml:",chardata"`
	Xsd            string   `xml:"xsd,attr,omitempty"`
	Xsi            string   `xml:"xsi,attr,omitempty"`
	ShipBlueprints struct {
		Text          string `xml:",chardata"`
		ShipBlueprint struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr,omitempty"`
			ID   struct {
				Text    string `xml:",chardata"`
				Type    string `xml:"Type,attr,omitempty"`
				Subtype string `xml:"Subtype,attr,omitempty"`
			} `xml:"Id,omitempty"`
			DisplayName string `xml:"DisplayName,omitempty"`
			CubeGrids   struct {
				Text     string `xml:",chardata"`
				CubeGrid struct {
					Text                   string `xml:",chardata"`
					SubtypeName            string `xml:"SubtypeName,omitempty"`
					EntityId               string `xml:"EntityId,omitempty"`
					PersistentFlags        string `xml:"PersistentFlags,omitempty"`
					PositionAndOrientation struct {
						Text     string `xml:",chardata"`
						Position struct {
							Text string `xml:",chardata"`
							X    string `xml:"x,attr,omitempty"`
							Y    string `xml:"y,attr,omitempty"`
							Z    string `xml:"z,attr,omitempty"`
						} `xml:"Position,omitempty"`
						Forward struct {
							Text string `xml:",chardata"`
							X    string `xml:"x,attr,omitempty"`
							Y    string `xml:"y,attr,omitempty"`
							Z    string `xml:"z,attr,omitempty"`
						} `xml:"Forward,omitempty"`
						Up struct {
							Text string `xml:",chardata"`
							X    string `xml:"x,attr,omitempty"`
							Y    string `xml:"y,attr,omitempty"`
							Z    string `xml:"z,attr,omitempty"`
						} `xml:"Up,omitempty"`
						Orientation struct {
							Text string `xml:",chardata"`
							X    string `xml:"X,omitempty"`
							Y    string `xml:"Y,omitempty"`
							Z    string `xml:"Z,omitempty"`
							W    string `xml:"W,omitempty"`
						} `xml:"Orientation,omitempty"`
					} `xml:"PositionAndOrientation,omitempty"`
					LocalPositionAndOrientation struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr,omitempty"`
					} `xml:"LocalPositionAndOrientation,omitempty"`
					GridSizeEnum string `xml:"GridSizeEnum,omitempty"`
					CubeBlocks   struct {
						Text                     string `xml:",chardata"`
						MyObjectBuilderCubeBlock []struct {
							Text        string `xml:",chardata"`
							Type        string `xml:"type,attr,omitempty"`
							SubtypeName string `xml:"SubtypeName,omitempty"`
							Min         struct {
								Text string `xml:",chardata"`
								X    string `xml:"x,attr,omitempty"`
								Y    string `xml:"y,attr,omitempty"`
								Z    string `xml:"z,attr,omitempty"`
							} `xml:"Min,omitempty"`
							ColorMaskHSV struct {
								Text string `xml:",chardata"`
								X    string `xml:"x,attr,omitempty"`
								Y    string `xml:"y,attr,omitempty"`
								Z    string `xml:"z,attr,omitempty"`
							} `xml:"ColorMaskHSV,omitempty"`
							SkinSubtypeId    string `xml:"SkinSubtypeId,omitempty"`
							BuiltBy          string `xml:"BuiltBy,omitempty"`
							EntityId         string `xml:"EntityId,omitempty"`
							BlockOrientation struct {
								Text    string `xml:",chardata"`
								Forward string `xml:"Forward,attr,omitempty"`
								Up      string `xml:"Up,attr,omitempty"`
							} `xml:"BlockOrientation,omitempty"`
							Owner               string `xml:"Owner,omitempty"`
							ShareMode           string `xml:"ShareMode,omitempty"`
							CustomName          string `xml:"CustomName,omitempty"`
							ShowOnHUD           string `xml:"ShowOnHUD,omitempty"`
							ShowInTerminal      string `xml:"ShowInTerminal,omitempty"`
							ShowInToolbarConfig string `xml:"ShowInToolbarConfig,omitempty"`
							ShowInInventory     string `xml:"ShowInInventory,omitempty"`
							NumberInGrid        string `xml:"NumberInGrid,omitempty"`
							Enabled             string `xml:"Enabled,omitempty"`
							ComponentContainer  struct {
								Text       string `xml:",chardata"`
								Components struct {
									Text          string `xml:",chardata"`
									ComponentData []struct {
										Text      string `xml:",chardata"`
										TypeId    string `xml:"TypeId,omitempty"`
										Component struct {
											Text                   string `xml:",chardata"`
											Type                   string `xml:"type,attr,omitempty"`
											Repeat                 string `xml:"Repeat,omitempty"`
											TimeToEvent            string `xml:"TimeToEvent,omitempty"`
											SetTimeMinutes         string `xml:"SetTimeMinutes,omitempty"`
											TimerEnabled           string `xml:"TimerEnabled,omitempty"`
											RemoveEntityOnTimer    string `xml:"RemoveEntityOnTimer,omitempty"`
											TimerType              string `xml:"TimerType,omitempty"`
											FramesFromLastTrigger  string `xml:"FramesFromLastTrigger,omitempty"`
											TimerTickInFrames      string `xml:"TimerTickInFrames,omitempty"`
											IsSessionUpdateEnabled string `xml:"IsSessionUpdateEnabled,omitempty"`
											CanPutItems            string `xml:"CanPutItems,omitempty"`
											Items                  string `xml:"Items,omitempty"`
											NextItemId             string `xml:"nextItemId,omitempty"`
											Volume                 string `xml:"Volume,omitempty"`
											Mass                   string `xml:"Mass,omitempty"`
											MaxItemCount           string `xml:"MaxItemCount,omitempty"`
											Size                   struct {
												Text string `xml:",chardata"`
												Nil  string `xml:"nil,attr,omitempty"`
											} `xml:"Size,omitempty"`
											InventoryFlags      string `xml:"InventoryFlags,omitempty"`
											RemoveEntityOnEmpty string `xml:"RemoveEntityOnEmpty,omitempty"`
										} `xml:"Component,omitempty"`
									} `xml:"ComponentData,omitempty"`
								} `xml:"Components,omitempty"`
							} `xml:"ComponentContainer,omitempty"`
							IsStockpiling          string `xml:"IsStockpiling,omitempty"`
							FilledRatio            string `xml:"FilledRatio,omitempty"`
							AutoRefill             string `xml:"AutoRefill,omitempty"`
							DeformationRatio       string `xml:"DeformationRatio,omitempty"`
							MasterToSlaveTransform struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr,omitempty"`
							} `xml:"MasterToSlaveTransform,omitempty"`
							MasterToSlaveGrid struct {
								Text string `xml:",chardata"`
								Nil  string `xml:"nil,attr,omitempty"`
							} `xml:"MasterToSlaveGrid,omitempty"`
							IsMaster                       string `xml:"IsMaster,omitempty"`
							TradingEnabled                 string `xml:"TradingEnabled,omitempty"`
							AutoUnlockTime                 string `xml:"AutoUnlockTime,omitempty"`
							TimeOfConnection               string `xml:"TimeOfConnection,omitempty"`
							IsParkingEnabled               string `xml:"IsParkingEnabled,omitempty"`
							IsPowerTransferOverrideEnabled string `xml:"IsPowerTransferOverrideEnabled,omitempty"`
							IsApproaching                  string `xml:"IsApproaching,omitempty"`
							IsConnecting                   string `xml:"IsConnecting,omitempty"`
						} `xml:"MyObjectBuilder_CubeBlock,omitempty"`
					} `xml:"CubeBlocks,omitempty"`
					XMirroxPlane struct {
						Text string `xml:",chardata"`
						X    string `xml:"x,attr,omitempty"`
						Y    string `xml:"y,attr,omitempty"`
						Z    string `xml:"z,attr,omitempty"`
					} `xml:"XMirroxPlane,omitempty"`
					DampenersEnabled string `xml:"DampenersEnabled,omitempty"`
					ConveyorLines    struct {
						Text                        string `xml:",chardata"`
						MyObjectBuilderConveyorLine []struct {
							Text          string `xml:",chardata"`
							StartPosition struct {
								Text string `xml:",chardata"`
								X    string `xml:"x,attr,omitempty"`
								Y    string `xml:"y,attr,omitempty"`
								Z    string `xml:"z,attr,omitempty"`
							} `xml:"StartPosition,omitempty"`
							StartDirection string `xml:"StartDirection,omitempty"`
							EndPosition    struct {
								Text string `xml:",chardata"`
								X    string `xml:"x,attr,omitempty"`
								Y    string `xml:"y,attr,omitempty"`
								Z    string `xml:"z,attr,omitempty"`
							} `xml:"EndPosition,omitempty"`
							EndDirection string `xml:"EndDirection,omitempty"`
							Sections     struct {
								Text    string `xml:",chardata"`
								Section []struct {
									Text      string `xml:",chardata"`
									Direction string `xml:"Direction,attr,omitempty"`
									Length    string `xml:"Length,attr,omitempty"`
								} `xml:"Section,omitempty"`
							} `xml:"Sections,omitempty"`
							ConveyorLineType string `xml:"ConveyorLineType,omitempty"`
						} `xml:"MyObjectBuilder_ConveyorLine,omitempty"`
					} `xml:"ConveyorLines,omitempty"`
					BlockGroups struct {
						Text                      string `xml:",chardata"`
						MyObjectBuilderBlockGroup []struct {
							Text   string `xml:",chardata"`
							Name   string `xml:"Name,omitempty"`
							Blocks struct {
								Text     string `xml:",chardata"`
								Vector3I []struct {
									Text string `xml:",chardata"`
									X    string `xml:"X,omitempty"`
									Y    string `xml:"Y,omitempty"`
									Z    string `xml:"Z,omitempty"`
								} `xml:"Vector3I,omitempty"`
							} `xml:"Blocks,omitempty"`
						} `xml:"MyObjectBuilder_BlockGroup,omitempty"`
					} `xml:"BlockGroups,omitempty"`
					DisplayName         string `xml:"DisplayName,omitempty"`
					DestructibleBlocks  string `xml:"DestructibleBlocks,omitempty"`
					IsRespawnGrid       string `xml:"IsRespawnGrid,omitempty"`
					LocalCoordSys       string `xml:"LocalCoordSys,omitempty"`
					TargetingTargets    string `xml:"TargetingTargets,omitempty"`
					NPCGridClaimElapsed struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr,omitempty"`
					} `xml:"NPCGridClaimElapsed,omitempty"`
				} `xml:"CubeGrid,omitempty"`
			} `xml:"CubeGrids,omitempty"`
			EnvironmentType string                  `xml:"EnvironmentType,omitempty"`
			WorkshopId      string                  `xml:"WorkshopId,omitempty"`
			WorkshopIds     []shared.WorkshopIDItem `xml:"WorkshopIds,omitempty"`
			OwnerSteamId    string                  `xml:"OwnerSteamId,omitempty"`
			Points          string                  `xml:"Points,omitempty"`
		} `xml:"ShipBlueprint,omitempty"`
	} `xml:"ShipBlueprints,omitempty"`
}
