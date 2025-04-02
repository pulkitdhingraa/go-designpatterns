package iterator

type Playlist struct {
	Songs []*Song
}

func (p *Playlist) NewIterator() Iterator{
	return &PlaylistIterator{playlist: p, index: 0}
}