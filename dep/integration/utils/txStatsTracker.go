package utils

import (
	"sort"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

// StatsConfig captures configuration details of an experiment
type StatsConfig struct {
	numCollNodes    int // number of collection nodes
	numConsNodes    int // number of consensus nodes
	numExecNodes    int // number of execution nodes
	numVerfNodes    int // number of verification nodes
	numCollClusters int // number of collection clusters
	txBatchSize     int // transaction batch size (tx sent at the same time)

}

// TxStats holds stats about execution of a transaction
type TxStats struct {
	TTF       time.Duration
	TTE       time.Duration
	TTS       time.Duration
	isExpired bool
}

// TxStatsTracker keep track of TxStats
type TxStatsTracker struct {
	config  *StatsConfig
	txStats []*TxStats
	mux     sync.Mutex
}

// NewTxStatsTracker returns a new instance of StatsTracker
func NewTxStatsTracker(config *StatsConfig) *TxStatsTracker {
	return &TxStatsTracker{
		config:  config,
		txStats: make([]*TxStats, 0),
	}
}

// AddTxStats adds a new TxStats to the tracker
func (st *TxStatsTracker) AddTxStats(tt *TxStats) {
	st.mux.Lock()
	defer st.mux.Unlock()
	st.txStats = append(st.txStats, tt)
}

// TotalTxSubmited returns the total transaction submited
func (st *TxStatsTracker) TotalTxSubmited() int {
	return len(st.txStats)
}

// TxFailureRate returns the number of expired transactions divided by total number of transactions
func (st *TxStatsTracker) TxFailureRate() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()
	counter := 0
	for _, t := range st.txStats {
		if t.isExpired {
			counter++
		}
	}
	return float64(counter) / float64(len(st.txStats))
}

// AvgTTF returns the average transaction time to finality (in seconds)
func (st *TxStatsTracker) AvgTTF() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()
	sum := float64(0)
	count := 0
	for _, t := range st.txStats {
		if !t.isExpired && t.TTF > 0 {
			sum += t.TTF.Seconds()
			count++
		}
	}
	return sum / float64(count)
}

// MedTTF returns the median transaction time to finality (in seconds)
func (st *TxStatsTracker) MedTTF() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()

	observations := make([]float64, 0)
	for _, t := range st.txStats {
		if !t.isExpired && t.TTF > 0 {
			observations = append(observations, t.TTF.Seconds())
		}
	}
	return computeMedian(observations)
}

// MaxTTF returns the maximum transaction time to finality (in seconds)
func (st *TxStatsTracker) MaxTTF() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()
	max := float64(0)
	for _, t := range st.txStats {
		if !t.isExpired && t.TTF > 0 {
			if t.TTF.Seconds() > max {
				max = t.TTF.Seconds()
			}
		}
	}
	return max
}

// AvgTTE returns the average transaction time to execution (in seconds)
func (st *TxStatsTracker) AvgTTE() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()
	sum := float64(0)
	count := 0
	for _, t := range st.txStats {
		if !t.isExpired && t.TTE > 0 {
			sum += t.TTE.Seconds()
			count++
		}
	}
	return sum / float64(count)
}

// MedTTE returns the median transaction time to execution (in seconds)
func (st *TxStatsTracker) MedTTE() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()

	observations := make([]float64, 0)
	for _, t := range st.txStats {
		if !t.isExpired && t.TTE > 0 {
			observations = append(observations, t.TTE.Seconds())
		}
	}
	return computeMedian(observations)
}

// MaxTTE returns the maximum transaction time to execution (in seconds)
func (st *TxStatsTracker) MaxTTE() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()
	max := float64(0)
	for _, t := range st.txStats {
		if !t.isExpired && t.TTE > 0 {
			if t.TTE.Seconds() > max {
				max = t.TTE.Seconds()
			}
		}
	}
	return max
}

// AvgTTS return the average transaction time to seal (in seconds)
func (st *TxStatsTracker) AvgTTS() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()
	sum := float64(0)
	count := 0
	for _, t := range st.txStats {
		if !t.isExpired && t.TTS > 0 {
			sum += t.TTS.Seconds()
			count++
		}
	}
	return sum / float64(count)
}

// MaxTTS returns the maximum transaction time to seal (in seconds)
func (st *TxStatsTracker) MaxTTS() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()
	max := float64(0)
	for _, t := range st.txStats {
		if !t.isExpired && t.TTS > 0 {
			if t.TTS.Seconds() > max {
				max = t.TTS.Seconds()
			}
		}
	}
	return max
}

// MedTTS returns the median transaction time to seal (in seconds)
func (st *TxStatsTracker) MedTTS() float64 {
	st.mux.Lock()
	defer st.mux.Unlock()

	observations := make([]float64, 0)
	for _, t := range st.txStats {
		if !t.isExpired && t.TTS > 0 {
			observations = append(observations, t.TTS.Seconds())
		}
	}
	return computeMedian(observations)
}

func (st *TxStatsTracker) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"index",
		"TTF (s)",
		"TTE (s)",
		"TTS (s)",
		"isExpired"})
	for i, tx := range st.txStats {
		t.AppendRow(table.Row{i,
			tx.TTF,
			tx.TTE,
			tx.TTS})

	}
	return t.Render()
}

func (st *TxStatsTracker) Digest() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"total TX",
		"failure rate",
		"avg TTF",
		"median TTF",
		"max TTF",
		"avg TTE",
		"median TTE",
		"max TTE",
		"avg TTS",
		"median TTS",
		"max TTS"})
	t.AppendRow(table.Row{st.TotalTxSubmited(),
		st.TxFailureRate(),
		st.AvgTTF(),
		st.MedTTF(),
		st.MaxTTF(),
		st.AvgTTE(),
		st.MedTTE(),
		st.MaxTTE(),
		st.AvgTTS(),
		st.MedTTS(),
		st.MaxTTS()})
	return t.Render()
}

func computeMedian(obs []float64) float64 {
	sort.Float64s(obs)
	switch len(obs) {
	case 0:
		return float64(0)
	case 1:
		return obs[0]
	case 2:
		return (obs[0] + obs[1]) / float64(2)
	default:
		if len(obs)%2 == 0 { // even
			return (obs[(len(obs)/2)-1] + obs[len(obs)/2]) / 2
		}
		return obs[len(obs)/2]
	}
}
